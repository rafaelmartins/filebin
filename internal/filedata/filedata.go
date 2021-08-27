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

func newfd(id string) (*FileData, error) {
	s, err := settings.Get()
	if err != nil {
		return nil, err
	}

	filename, mimetype, size, timestamp, err := s.Backend.ReadMetadata(id)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	fd := &FileData{
		id:        id,
		Filename:  filename,
		Mimetype:  mimetype,
		Size:      size,
		Timestamp: timestamp,
	}

	reg.m.Lock()
	reg.data[fd.id] = fd
	reg.dataslice = append(reg.dataslice, fd)
	reg.m.Unlock()

	return fd, nil
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
		if _, err := newfd(id); err != nil {
			return err
		}
	}

	reg.m.Lock()
	defer reg.m.Unlock()

	sort.Sort(&byDate{reg.dataslice})

	return nil
}

func processFile(fh *multipart.FileHeader) (*FileData, error) {
	s, err := settings.Get()
	if err != nil {
		return nil, err
	}

	f, err := fh.Open()
	defer f.Close()

	if fh.Size > int64(s.UploadMaxSizeMb)*1024*1024 {
		return nil, errors.New("filedata: uploaded file bigger than allowed size")
	}

	m, err := mime.Detect(f, fh)
	if err != nil {
		return nil, err
	}

	if _, err := f.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	var (
		fid string
		n   int64
	)

	for {
		var err error
		fid, err = id.Generate(s.IdLength)
		if err != nil {
			return nil, err
		}

		n, err = s.Backend.Write(fid, f, fh.Filename, m)
		if err != nil {
			if !os.IsExist(err) {
				return nil, err
			}
			if _, err := f.Seek(0, io.SeekStart); err != nil {
				return nil, err
			}
		} else {
			break
		}
	}

	if n != fh.Size {
		s.Backend.Delete(fid)
		if n < fh.Size {
			return nil, errors.New("filedata: write: unexpected eof")
		}
		return nil, errors.New("filedata: write: mismatched file size")
	}

	return newfd(fid)
}

func NewFromRequest(r *http.Request) ([]*FileData, error) {
	if r == nil {
		return nil, errors.New("filedata: nil request")
	}

	if err := r.ParseMultipartForm(32 * 1024 * 1024); err != nil {
		return nil, err
	}

	if r.MultipartForm == nil || r.MultipartForm.File == nil {
		return nil, errors.New("filedata: no files")
	}
	defer r.MultipartForm.RemoveAll()

	fhs := r.MultipartForm.File["file"]
	if len(fhs) == 0 {
		return nil, errors.New("filedata: no files")
	}

	fds := []*FileData{}
	errl := []string{}
	for i, fh := range fhs {
		fd, err := processFile(fh)
		if err != nil {
			fds = append(fds, nil)
			errl = append(errl, fmt.Sprintf("%d: %s", i, err.Error()))
			continue
		}
		fds = append(fds, fd)
	}

	var err error
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

	return s.Backend.Delete(fd.id)
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

func (f *FileData) Serve(w http.ResponseWriter, r *http.Request, filename string, mimetype string, attachment bool) error {
	s, err := settings.Get()
	if err != nil {
		return err
	}

	return s.Backend.Serve(w, r, f.id, filename, mimetype, attachment)
}

func (f *FileData) Read() (io.ReadCloser, error) {
	s, err := settings.Get()
	if err != nil {
		return nil, err
	}

	return s.Backend.Read(f.id)
}
