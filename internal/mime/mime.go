package mime

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"

	"github.com/danwakefield/fnmatch"
	"github.com/rafaelmartins/filebin/internal/mime/magic"
)

var (
	contentType = textproto.CanonicalMIMEHeaderKey("content-type")
	errNotFound = errors.New("mime: type not found")
)

func detectFromData(r io.Reader) (string, error) {
	// read start of file and rewind
	var buf [512]byte
	n, _ := io.ReadFull(r, buf[:])

	// magic is usually the most reliable tool, let's try it first
	if m, err := magic.Detect(buf[:n]); err == nil && m != "" {
		if v, ok := magicMap[m]; ok && v != "" {
			return v, nil
		}
		return m, nil
	}

	// then try signature detection from net/http
	if m := http.DetectContentType(buf[:n]); m != "" {
		return m, nil
	}

	return "", errNotFound
}

func detectFromFilename(filename string) (string, error) {
	index := -1
	var match *mimeType
	for _, t := range registry {
		for i, p := range t.patterns {
			if fnmatch.Match(p, filename, 0) {
				if i == 0 {
					return t.name, nil
				}
				if index == -1 || i < index {
					index = i
					match = t
				}
				break
			}
		}
	}
	if match != nil && index != -1 {
		return match.name, nil
	}
	return "", errNotFound
}

func Detect(f io.Reader, fh *multipart.FileHeader) (string, error) {
	if f == nil || fh == nil {
		return "", errNotFound
	}

	if fh.Filename != "-" {
		if m, err := detectFromFilename(fh.Filename); err == nil && m != "" {
			return m, nil
		}
	}

	if m, err := detectFromData(f); err == nil && m != "" {
		return m, nil
	}

	// our last resource is trusting the mime type sent by http client
	// this is usually good enough for browsers, but not enough for curl
	for key, l := range fh.Header {
		if len(l) > 0 && contentType == key {
			return l[0], nil
		}
	}

	return "", errNotFound
}

func GetExtension(mimetype string, filename string) string {
	fn := filepath.Base(filename)
	for _, m := range registry {
		if m.name == mimetype {
			if len(m.patterns) == 0 {
				continue
			}
			for _, p := range m.patterns {
				if fn == p {
					return ""
				}
			}
			for _, p := range m.patterns {
				if ext := filepath.Ext(p); ext != "" && !strings.Contains(ext, "*") {
					return ext
				}
			}
		}
	}
	return ""
}
