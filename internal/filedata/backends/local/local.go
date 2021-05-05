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
	"time"
)

type Local struct {
	dir string
}

type metadata struct {
	Filename  string    `json:"filename"`
	Mimetype  string    `json:"mimetype"`
	Timestamp time.Time `json:"timestamp"`
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

func (l *Local) Read(id string) (io.ReadCloser, error) {
	return os.Open(filepath.Join(l.dir, id))
}

func (l *Local) ReadMetadata(id string) (string, string, int64, time.Time, error) {
	fn := filepath.Join(l.dir, id+".json")
	fp, err := os.Open(fn)
	if err != nil {
		return "", "", 0, time.Time{}, err
	}
	defer fp.Close()

	v := &metadata{}
	if err := json.NewDecoder(fp).Decode(v); err != nil {
		return "", "", 0, time.Time{}, err
	}

	st, err := os.Stat(filepath.Join(l.dir, id))
	if err != nil {
		return "", "", 0, time.Time{}, err
	}
	return v.Filename, v.Mimetype, st.Size(), v.Timestamp, nil
}

func (l *Local) writeJSON(id string, v *metadata) error {
	fn := filepath.Join(l.dir, id+".json")
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		return err
	}
	defer fp.Close()
	return json.NewEncoder(fp).Encode(v)
}

func (l *Local) deleteJSON(id string) error {
	return os.Remove(filepath.Join(l.dir, id+".json"))
}

func (l *Local) Write(id string, r io.ReadSeeker, filename string, mimetype string) (int64, error) {
	v := &metadata{
		Filename:  filename,
		Mimetype:  mimetype,
		Timestamp: time.Now().UTC(),
	}
	if err := l.writeJSON(id, v); err != nil {
		return 0, err
	}

	fn := filepath.Join(l.dir, id)
	fp, err := os.OpenFile(fn, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		l.deleteJSON(id)
		return 0, err
	}
	defer fp.Close()
	return io.Copy(fp, r)
}

func (l *Local) Delete(id string) error {
	err1 := l.deleteJSON(id)
	err2 := os.Remove(filepath.Join(l.dir, id))
	if err1 != nil && err2 != nil {
		return errors.New(err1.Error() + " | " + err2.Error())
	}
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func (l *Local) Serve(w http.ResponseWriter, r *http.Request, id string, filename string, mimetype string, attachment bool) error {
	fn := filepath.Join(l.dir, id)
	w.Header().Set("Content-Type", mimetype)
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
