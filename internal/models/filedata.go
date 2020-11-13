package models

import (
	"errors"
	"path/filepath"
	"time"
)

var (
	ErrFileDataDirNotDefined = errors.New("models: filedata: directory not defined")
)

type FileData struct {
	Filename  string    `json:"filename"`
	Mimetype  string    `json:"mimetype"`
	Size      int64     `json:"size"`
	Timestamp time.Time `json:"timestamp"`
}

type FileDataModel struct {
	Dir string
}

func (f *FileDataModel) Read(id string, fd *FileData) error {
	if f.Dir == "" {
		return ErrFileDataDirNotDefined
	}
	return modelRead(filepath.Join(f.Dir, id+".json"), fd)
}

func (f *FileDataModel) Write(id string, fd *FileData) error {
	if f.Dir == "" {
		return ErrFileDataDirNotDefined
	}
	return modelWrite(filepath.Join(f.Dir, id+".json"), fd)
}

func (f *FileDataModel) GetFilename(id string) string {
	if f.Dir == "" {
		return ""
	}
	return filepath.Join(f.Dir, id)
}
