package settings

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rafaelmartins/filebin/internal/filedata/backends"
	"github.com/rafaelmartins/filebin/internal/filedata/backends/s3"
)

var (
	settings *Settings
)

type Settings struct {
	AuthRealm       string
	AuthUsername    string
	AuthPassword    string
	BaseUrl         string
	HighlightStyle  string
	IdLength        uint8
	ListenAddr      string
	UploadMaxSizeMb uint
	IndexFooter     string

	S3Options  s3.S3Options
	StorageDir string

	Backend backends.Backend
}

func getString(key string, def string, required bool) (string, error) {
	if v, found := os.LookupEnv(key); found {
		if required && v == "" {
			return "", fmt.Errorf("settings: %s empty", key)
		}
		return v, nil
	}
	if required && def == "" {
		return "", fmt.Errorf("settings: %s missing", key)
	}
	return def, nil
}

func getUint(key string, def uint64, required bool, base int, bitSize int) (uint64, error) {
	v, err := getString(key, strconv.FormatUint(def, base), required)
	if err != nil {
		return 0, err
	}
	v2, err := strconv.ParseUint(v, base, bitSize)
	if err != nil {
		return 0, err
	}
	if required && v2 == 0 {
		return 0, fmt.Errorf("settings: %s empty", key)
	}
	return v2, nil
}

func getBool(key string, def bool) (bool, error) {
	v, err := getString(key, strconv.FormatBool(def), true)
	if err != nil {
		return false, err
	}
	v2, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}
	return v2, nil
}

func Get() (*Settings, error) {
	if settings != nil {
		return settings, nil
	}

	var err error
	s := &Settings{}

	s.AuthRealm, err = getString("FILEBIN_AUTH_REALM", "filebin", true)
	if err != nil {
		return nil, err
	}

	s.AuthUsername, err = getString("FILEBIN_AUTH_USERNAME", "", true)
	if err != nil {
		return nil, err
	}

	s.AuthPassword, err = getString("FILEBIN_AUTH_PASSWORD", "", true)
	if err != nil {
		return nil, err
	}

	s.BaseUrl, err = getString("FILEBIN_BASE_URL", "", false)
	if err != nil {
		return nil, err
	}

	s.HighlightStyle, err = getString("FILEBIN_HIGHLIGHT_STYLE", "trac", true)
	if err != nil {
		return nil, err
	}

	idLength, err := getUint("FILEBIN_ID_LENGTH", 8, true, 10, 8)
	if err != nil {
		return nil, err
	}
	if idLength < 8 {
		return nil, errors.New("FILEBIN_ID_LENGTH must be >= 8")
	}
	s.IdLength = uint8(idLength)

	s.ListenAddr, err = getString("FILEBIN_LISTEN_ADDR", ":8000", true)
	if err != nil {
		return nil, err
	}

	s.S3Options.AccessKeyId, err = getString("FILEBIN_S3_ACCESS_KEY_ID", "", false)
	if err != nil {
		return nil, err
	}

	s.S3Options.SecretAccessKey, err = getString("FILEBIN_S3_SECRET_ACCESS_KEY", "", false)
	if err != nil {
		return nil, err
	}

	s.S3Options.Endpoint, err = getString("FILEBIN_S3_ENDPOINT", "", false)
	if err != nil {
		return nil, err
	}

	s.S3Options.Region, err = getString("FILEBIN_S3_REGION", "", false)
	if err != nil {
		return nil, err
	}

	s.S3Options.Bucket, err = getString("FILEBIN_S3_BUCKET", "", false)
	if err != nil {
		return nil, err
	}

	s3PresignExpireMinutes, err := getUint("FILEBIN_S3_PRESIGN_EXPIRE_MINUTES", 5, true, 10, 0)
	if err != nil {
		return nil, err
	}
	s.S3Options.PresignExpire = time.Duration(s3PresignExpireMinutes) * time.Minute

	s.S3Options.ProxyData, err = getBool("FILEBIN_S3_PROXY_DATA", false)
	if err != nil {
		return nil, err
	}

	s.S3Options.ForcePathStyle, err = getBool("FILEBIN_S3_FORCE_PATH_STYLE", false)
	if err != nil {
		return nil, err
	}

	s.S3Options.SslInsecure, err = getBool("FILEBIN_S3_SSL_INSECURE", false)
	if err != nil {
		return nil, err
	}

	s.S3Options.SslCertificate, err = getString("FILEBIN_S3_SSL_CERTIFICATE", "", false)
	if err != nil {
		return nil, err
	}

	s.StorageDir, err = getString("FILEBIN_STORAGE_DIR", "", false)
	if err != nil {
		return nil, err
	}

	s.IndexFooter, err = getString("FILEBIN_INDEX_FOOTER", "", false)
	if err != nil {
		return nil, err
	}

	uploadMaxSizeMb, err := getUint("FILEBIN_UPLOAD_MAX_SIZE_MB", 10, true, 10, 0)
	if err != nil {
		return nil, err
	}
	if uploadMaxSizeMb == 0 {
		return nil, errors.New("FILEBIN_UPLOAD_MAX_SIZE_MB must be > 0")
	}
	s.UploadMaxSizeMb = uint(uploadMaxSizeMb)

	s.Backend, err = backends.Lookup(s.StorageDir, s.S3Options)
	if err != nil {
		return nil, err
	}

	settings = s

	return s, nil
}
