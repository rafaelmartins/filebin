package filedata

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/rafaelmartins/filebin/internal/id"
	"github.com/rafaelmartins/filebin/internal/mime"
	"github.com/rafaelmartins/filebin/internal/settings"
)

var (
	ErrNotFound = errors.New("filedata: not found")

	reg = &registry{data: map[string]*FileData{}}
)

type registry struct {
	data      map[string]*FileData
	dataslice []*FileData
	m         sync.RWMutex
}

type FileData struct {
	id        string
	Filename  string    `json:"filename"`
	Mimetype  string    `json:"mimetype"`
	Size      int64     `json:"size"`
	Timestamp time.Time `json:"timestamp"`
}

type byDate struct {
	data []*FileData
}

func (b *byDate) Len() int {
	return len(b.data)
}

func (b *byDate) Swap(i int, j int) {
	b.data[i], b.data[j] = b.data[j], b.data[i]
}

func (b *byDate) Less(i int, j int) bool {
	return b.data[i].Timestamp.UnixNano() < b.data[j].Timestamp.UnixNano()
}

func Init() error {
	s, err := settings.Get()
	if err != nil {
		return err
	}

	ids, err := s.Backend.List()
	if err != nil {
		return err
	}

	for _, id := range ids {
		fd := &FileData{
			id: id,
		}

		if err := s.Backend.ReadJSON(fd.id, fd); err != nil {
			if os.IsNotExist(err) {
				return ErrNotFound
			}
			return err
		}

		reg.m.Lock()
		reg.data[fd.id] = fd
		reg.dataslice = append(reg.dataslice, fd)
		reg.m.Unlock()
	}

	reg.m.Lock()
	defer reg.m.Unlock()

	sort.Sort(&byDate{reg.dataslice})

	return nil
}

func processFile(fh *multipart.FileHeader, size int64, idLength uint8) (*FileData, error) {
	s, err := settings.Get()
	if err != nil {
		return nil, err
	}

	f, err := fh.Open()
	defer f.Close()

	if fh.Size > size {
		return nil, errors.New("filedata: uploaded file bigger than allowed size")
	}

	m, err := mime.Detect(f, fh)
	if err != nil {
		return nil, err
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	fd := &FileData{
		Filename:  fh.Filename,
		Mimetype:  m,
		Size:      fh.Size,
		Timestamp: time.Now().UTC(),
	}

	for {
		var err error
		fd.id, err = id.Generate(idLength)
		if err != nil {
			return nil, err
		}

		if err := s.Backend.WriteJSON(fd.id, fd); err != nil {
			if !os.IsExist(err) {
				return nil, err
			}
		} else {
			break
		}
	}

	n, err := s.Backend.WriteData(fd.id, f)
	if err != nil {
		s.Backend.DeleteData(fd.id)
		return nil, err
	}
	if n != fh.Size {
		s.Backend.DeleteData(fd.id)
		if n < fh.Size {
			return nil, errors.New("filedata: write: unexpected eof")
		}
		return nil, errors.New("filedata: write: mismatched file size")
	}

	reg.m.Lock()
	defer reg.m.Unlock()
	reg.data[fd.id] = fd
	reg.dataslice = append(reg.dataslice, fd)

	return fd, nil
}

func NewFromRequest(r *http.Request) ([]*FileData, error) {
	if r == nil {
		return nil, errors.New("filedata: nil request")
	}

	s, err := settings.Get()
	if err != nil {
		return nil, err
	}

	if err := r.ParseMultipartForm(32 * 1024 * 1024); err != nil {
		return nil, err
	}

	if r.MultipartForm == nil || r.MultipartForm.File == nil {
		return nil, errors.New("filedata: no files")
	}

	fhs := r.MultipartForm.File["file"]
	if len(fhs) == 0 {
		return nil, errors.New("filedata: no files")
	}

	size := int64(s.UploadMaxSizeMb) * 1024 * 1024

	fds := []*FileData{}
	errl := []string{}
	for i, fh := range fhs {
		fd, err := processFile(fh, size, s.IdLength)
		if err != nil {
			fds = append(fds, nil)
			errl = append(errl, fmt.Sprintf("%d: %s", i, err.Error()))
			continue
		}
		fds = append(fds, fd)
	}

	if len(errl) > 0 {
		err = errors.New(strings.Join(errl, " | "))
	}

	if len(fds) == 0 {
		return nil, err
	}
	return fds, err
}

func NewFromId(id string) (*FileData, error) {
	reg.m.RLock()
	defer reg.m.RUnlock()

	if fd, ok := reg.data[id]; ok {
		return fd, nil
	}

	return nil, ErrNotFound
}

func ForEach(f func(*FileData)) {
	reg.m.RLock()
	defer reg.m.RUnlock()

	for _, fd := range reg.dataslice {
		f(fd)
	}
}

func Delete(id string) error {
	s, err := settings.Get()
	if err != nil {
		return err
	}

	fd, err := NewFromId(id)
	if err != nil {
		return err
	}

	reg.m.Lock()
	defer reg.m.Unlock()

	delete(reg.data, fd.id)

	n := []*FileData{}
	for _, v := range reg.dataslice {
		if v.id != id {
			n = append(n, v)
		}
	}
	reg.dataslice = n

	if err := s.Backend.DeleteJSON(fd.id); err != nil {
		return err
	}

	return s.Backend.DeleteData(fd.id)
}

func (f *FileData) GetId() string {
	return f.id
}

func (f *FileData) GetFilename() string {
	if filepath.Ext(f.Filename) == "" {
		fn := f.Filename

		// exceptions
		// FIXME: handle other filenames with just special chars
		if fn == "-" {
			fn = "stdin"
		}

		// try to find a file extension
		fn += mime.GetExtension(f.Mimetype, f.Filename)
		return fn
	}
	return f.Filename
}

func (f *FileData) ServeData(w http.ResponseWriter, r *http.Request, contentType string, filename string, attachment bool) error {
	s, err := settings.Get()
	if err != nil {
		return err
	}

	return s.Backend.ServeData(w, r, f.id, contentType, filename, attachment)
}

func (f *FileData) OpenData() (io.ReadCloser, error) {
	s, err := settings.Get()
	if err != nil {
		return nil, err
	}

	return s.Backend.OpenData(f.id)
}
