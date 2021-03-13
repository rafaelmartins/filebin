package local

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

type Local struct {
	dir string
}

func NewLocal(dir string) (*Local, error) {
	st, err := os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		if err := os.MkdirAll(dir, 0777); err != nil {
			return nil, err
		}
	}
	if !st.IsDir() {
		return nil, errors.New("local: defined storage directory is not a directory")
	}
	return &Local{dir: dir}, nil
}

func (l *Local) Name() string {
	return "Local"
}

func (l *Local) List() ([]string, error) {
	files, err := ioutil.ReadDir(l.dir)
	if err != nil {
		return nil, err
	}

	rv := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if fn := file.Name(); filepath.Ext(fn) == ".json" {
			rv = append(rv, fn[:len(fn)-5])
		}
	}
	return rv, nil
}

func (l *Local) ReadJSON(id string, v interface{}) error {
	fn := filepath.Join(l.dir, id+".json")
	fp, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer fp.Close()
	return json.NewDecoder(fp).Decode(v)
}

func (l *Local) WriteJSON(id string, v interface{}) error {
	fn := filepath.Join(l.dir, id+".json")
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	return json.NewEncoder(fp).Encode(v)
}

func (l *Local) DeleteJSON(id string) error {
	return os.Remove(filepath.Join(l.dir, id+".json"))
}

func (l *Local) OpenData(id string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(l.dir, id))
}

func (l *Local) WriteData(id string, r io.ReadSeeker) (int64, error) {
	fn := filepath.Join(l.dir, id)
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return 0, err
	}
	defer fp.Close()
	return io.Copy(fp, r)
}

func (l *Local) DeleteData(id string) error {
	return os.Remove(filepath.Join(l.dir, id))
}

func (l *Local) ServeData(w http.ResponseWriter, r *http.Request, id string, contentType string, filename string, attachment bool) error {
	fn := filepath.Join(l.dir, id)
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if filename != "" {
		if attachment {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
		} else {
			w.Header().Set("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, filename))
		}
	}
	http.ServeFile(w, r, fn)
	return nil
}
