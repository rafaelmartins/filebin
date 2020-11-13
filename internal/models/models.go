package models

import (
	"encoding/json"
	"errors"
	"os"
)

var (
	ErrDuplicated = errors.New("models: duplicated")
	ErrNotFound   = errors.New("models: not found")
)

func modelRead(filename string, obj interface{}) error {
	f, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotFound
		}
		return err
	}
	defer f.Close()

	return json.NewDecoder(f).Decode(obj)
}

func modelWrite(filename string, obj interface{}) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		if os.IsExist(err) {
			return ErrDuplicated
		}
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(obj)
}
