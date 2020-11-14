package filedata

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/alecthomas/chroma"
	"github.com/rafaelmartins/filebin/internal/id"
	"github.com/rafaelmartins/filebin/internal/mime"
	"github.com/rafaelmartins/filebin/internal/settings"
)

var (
	ErrDuplicated = errors.New("models: duplicated")
	ErrNotFound   = errors.New("models: not found")
)

type FileData struct {
	id        string
	Filename  string    `json:"filename"`
	Mimetype  string    `json:"mimetype"`
	Size      int64     `json:"size"`
	Timestamp time.Time `json:"timestamp"`
	lexer     chroma.Lexer
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

	if fh.Size > size {
		return nil, errors.New("filedata: uploaded file bigger than allowed size")
	}

	m, err := mime.Detect(f, fh)
	if err != nil {
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
		if err := fd.write(); err != nil {
			if err != ErrDuplicated {
				return nil, err
			}
		} else {
			break
		}
	}

	if err := fd.writeFile(f); err != nil {
		return nil, err
	}

	return fd, nil
}

func NewFromId(id string) (*FileData, error) {
	fd := &FileData{
		id: id,
	}
	if err := fd.read(); err != nil {
		return nil, err
	}
	return fd, nil
}

func (f *FileData) getFilename() (string, error) {
	if f.id == "" {
		return "", errors.New("filedata: id missing")
	}
	s, err := settings.Get()
	if err != nil {
		return "", err
	}
	return filepath.Join(s.StorageDir, f.id), nil
}

func (f *FileData) getJsonFilename() (string, error) {
	fn, err := f.getFilename()
	if err != nil {
		return "", err
	}
	return fn + ".json", nil
}

func (f *FileData) read() error {
	fn, err := f.getJsonFilename()
	if err != nil {
		return err
	}

	fp, err := os.Open(fn)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotFound
		}
		return err
	}
	defer fp.Close()

	return json.NewDecoder(fp).Decode(f)
}

func (f *FileData) write() error {
	fn, err := f.getJsonFilename()
	if err != nil {
		return err
	}

	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return ErrDuplicated
		}
		return err
	}
	defer fp.Close()

	return json.NewEncoder(fp).Encode(f)
}

func (f *FileData) readFile() ([]byte, error) {
	fn, err := f.getFilename()
	if err != nil {
		return nil, err
	}
	return ioutil.ReadFile(fn)
}

func (f *FileData) writeFile(d io.Reader) error {
	fn, err := f.getFilename()
	if err != nil {
		return err
	}

	fp, err := os.Create(fn)
	if err != nil {
		return err
	}

	n, err := io.Copy(fp, d)
	if err != nil {
		fp.Close()
		os.RemoveAll(fn)
		return err
	}
	if n != f.Size {
		fp.Close()
		os.RemoveAll(fn)
		if n < f.Size {
			return errors.New("filedata: write: unexpected eof")
		}
		return errors.New("filedata: write: mismatched file size")
	}

	return fp.Close()
}

func (f *FileData) ServeFile(w http.ResponseWriter, r *http.Request) error {
	fn, err := f.getFilename()
	if err != nil {
		return err
	}
	http.ServeFile(w, r, fn)
	return nil
}

func (f *FileData) GetId() string {
	return f.id
}
