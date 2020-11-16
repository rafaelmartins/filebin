package mime

import (
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"
	"strings"

	"github.com/rafaelmartins/filebin/internal/highlight"
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
		return m, nil
	}

	// then try signature detection from net/http
	if m := http.DetectContentType(buf[:n]); m != "" {
		return m, nil
	}

	return "", errNotFound
}

func detectFromFilename(filename string) (string, error) {
	// try source types first
	if m, err := highlight.DetectMimetype(filename); err == nil && m != "" {
		return m, nil
	}

	// then try from OS mime database (at least on unix)
	if m := mime.TypeByExtension(filepath.Ext(filename)); m != "" {
		return m, nil
	}

	return "", errNotFound
}

func Detect(f io.Reader, fh *multipart.FileHeader) (string, error) {
	if f == nil || fh == nil {
		return "", errNotFound
	}

	if filepath.Ext(fh.Filename) != "" {
		// files with extension can be easily matched without looking at the content
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

func GetExtension(mimetype string) string {
	// exceptions
	if strings.HasPrefix(mimetype, "text/plain") {
		return ".txt"
	}

	ext, err := mime.ExtensionsByType(mimetype)
	if err != nil {
		return ""
	}
	if len(ext) == 0 {
		return ""
	}
	return ext[0]
}
