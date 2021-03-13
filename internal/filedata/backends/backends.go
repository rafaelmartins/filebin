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
	ReadJSON(id string, v interface{}) error
	WriteJSON(id string, v interface{}) error
	DeleteJSON(id string) error
	OpenData(id string) (io.ReadCloser, error)
	WriteData(id string, r io.ReadSeeker) (int64, error)
	DeleteData(id string) error
	ServeData(w http.ResponseWriter, r *http.Request, id string, contentType string, filename string, attachment bool) error
}

func Lookup(dir string, s3AccessKeyId string, s3SecretAccessKey string, s3Endpoint string, s3Region string, s3Bucket string, s3PresignExpire time.Duration) (Backend, error) {
	if s3AccessKeyId != "" && s3SecretAccessKey != "" && s3Region != "" && s3Bucket != "" {
		return s3.NewS3(s3AccessKeyId, s3SecretAccessKey, s3Endpoint, s3Region, s3Bucket, s3PresignExpire), nil
	}

	if dir != "" {
		return local.NewLocal(dir)
	}

	return nil, errors.New("filedata: no data backend configured")
}
