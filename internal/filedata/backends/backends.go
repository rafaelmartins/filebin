package backends

import (
	"errors"
	"io"
	"net/http"
	"time"

	"github.com/rafaelmartins/filebin/internal/filedata/backends/local"
	"github.com/rafaelmartins/filebin/internal/filedata/backends/s3"
)

type Backend interface {
	Name() string
	List() ([]string, error)
	Read(id string) (io.ReadCloser, error)
	ReadMetadata(id string) (string, string, int64, time.Time, error)
	Write(id string, r io.ReadSeeker, filename string, mimetype string) (int64, error)
	Delete(id string) error
	Serve(w http.ResponseWriter, r *http.Request, id string, filename string, mimetype string, timestamp time.Time, attachment bool) error
}

func Lookup(dir string, s3Options s3.S3Options) (Backend, error) {
	if s3Options.AccessKeyId != "" && s3Options.SecretAccessKey != "" && s3Options.Region != "" && s3Options.Bucket != "" {
		return s3.NewS3(s3Options)
	}

	if dir != "" {
		return local.NewLocal(dir)
	}

	return nil, errors.New("filedata: no data backend configured")
}
