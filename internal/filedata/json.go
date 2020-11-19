package filedata

import (
	"encoding/json"
	"os"
)

func (f *FileData) getFilenameJSON() (string, error) {
	fn, err := f.getFilenameData()
	if err != nil {
		return "", err
	}
	return fn + ".json", nil
}

func (f *FileData) readJSON() error {
	fn, err := f.getFilenameJSON()
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

func (f *FileData) writeJSON() error {
	fn, err := f.getFilenameJSON()
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

func (f *FileData) deleteJSON() error {
	fn, err := f.getFilenameJSON()
	if err != nil {
		return err
	}
	return os.Remove(fn)
}
