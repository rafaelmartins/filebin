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
	Serve(w http.ResponseWriter, r *http.Request, id string, filename string, mimetype string, attachment bool) error
}

func Lookup(dir string, s3AccessKeyId string, s3SecretAccessKey string, s3Endpoint string, s3Region string, s3Bucket string, s3PresignExpire time.Duration, s3ProxyData bool) (Backend, error) {
	if s3AccessKeyId != "" && s3SecretAccessKey != "" && s3Region != "" && s3Bucket != "" {
		return s3.NewS3(
			s3AccessKeyId,
			s3SecretAccessKey,
			s3Endpoint,
			s3Region,
			s3Bucket,
			s3PresignExpire,
			s3ProxyData,
		), nil
	}

	if dir != "" {
		return local.NewLocal(dir)
	}

	return nil, errors.New("filedata: no data backend configured")
}
