package filedata

import (
	"errors"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/rafaelmartins/filebin/internal/id"
	"github.com/rafaelmartins/filebin/internal/mime"
	"github.com/rafaelmartins/filebin/internal/settings"
)

var (
	ErrDuplicated = errors.New("models: duplicated")
	ErrNotFound   = errors.New("models: not found")

	reg = &registry{data: map[string]*FileData{}}
)

type registry struct {
	data map[string]*FileData
	m    sync.Mutex
}

type FileData struct {
	id        string
	Filename  string    `json:"filename"`
	Mimetype  string    `json:"mimetype"`
	Size      int64     `json:"size"`
	Timestamp time.Time `json:"timestamp"`
	lexer     *string
}

func NewFromRequest(r *http.Request) (*FileData, error) {
	if r == nil {
		return nil, errors.New("filedata: nil request")
	}

	s, err := settings.Get()
	if err != nil {
		return nil, err
	}

	// we store all file data in memory, instead of letting the library use temp files.
	// this is intended to be used as a private service, it should not be an issue.
	size := int64(s.UploadMaxSizeMb) * 1024 * 1024
	if err := r.ParseMultipartForm(size); err != nil {
		return nil, err
	}

	f, fh, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
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
		fd.id, err = id.Generate(s.IdLength)
		if err != nil {
			return nil, err
		}
		if err := fd.writeJSON(); err != nil {
			if err != ErrDuplicated {
				return nil, err
			}
		} else {
			break
		}
	}

	if err := fd.writeData(f); err != nil {
		return nil, err
	}

	reg.m.Lock()
	defer reg.m.Unlock()
	reg.data[fd.id] = fd

	return fd, nil
}

func NewFromId(id string) (*FileData, error) {
	if fd, ok := reg.data[id]; ok {
		return fd, nil
	}

	fd := &FileData{
		id: id,
	}
	if err := fd.readJSON(); err != nil {
		return nil, err
	}

	reg.m.Lock()
	defer reg.m.Unlock()
	reg.data[fd.id] = fd

	return fd, nil
}

func (f *FileData) GetId() string {
	return f.id
}
