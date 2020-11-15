package filedata

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/rafaelmartins/filebin/internal/settings"
)

func (f *FileData) getFilenameData() (string, error) {
	if f.id == "" {
		return "", errors.New("filedata: id missing")
	}
	s, err := settings.Get()
	if err != nil {
		return "", err
	}
	return filepath.Join(s.StorageDir, f.id), nil
}

func (f *FileData) writeData(d io.Reader) error {
	fn, err := f.getFilenameData()
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

func (f *FileData) ServeData(w http.ResponseWriter, r *http.Request) error {
	fn, err := f.getFilenameData()
	if err != nil {
		return err
	}
	http.ServeFile(w, r, fn)
	return nil
}
