package upload

import (
	"errors"
	"io"
	"mime"
	"net/http"
	"net/textproto"
	"path/filepath"

	"github.com/rafaelmartins/filebin/internal/highlight"
	"github.com/rafaelmartins/filebin/internal/magic"
)

var (
	contentType     = textproto.CanonicalMIMEHeaderKey("content-type")
	errMimeNotFound = errors.New("upload: mime type not found")
)

func detectFromContent(r io.ReadSeeker) (string, error) {
	// read start of file and rewind
	var buf [512]byte
	n, _ := io.ReadFull(r, buf[:])
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	// magic is usually the most reliable tool, let's try it first
	if m, err := magic.Detect(buf[:n]); err == nil && m != "" {
		return m, nil
	}

	// then try signature detection from net/http
	if m := http.DetectContentType(buf[:n]); m != "" {
		return m, nil
	}

	return "", errMimeNotFound
}

func detectFromFilename(filename string) (string, error) {
	// try source types first
	if m := highlight.GetMimeType(filename); m != "" {
		return m, nil
	}

	// then try from OS mime database (at least on unix)
	if m := mime.TypeByExtension(filepath.Ext(filename)); m != "" {
		return m, nil
	}

	return "", errMimeNotFound
}

func getMimeType(filename string, r io.ReadSeeker, header textproto.MIMEHeader) (string, error) {
	if filepath.Ext(filename) != "" {
		// files with extension can be easily matched without looking at the content
		if m, err := detectFromFilename(filename); err == nil && m != "" {
			return m, nil
		}
	}

	if m, err := detectFromContent(r); err == nil && m != "" {
		return m, nil
	}

	// our last resource is trusting the mime type sent by http client
	// this is usually good enough for browsers, but not enough for curl
	for key, l := range header {
		if len(l) > 0 && contentType == key {
			return l[0], nil
		}
	}

	return "", errMimeNotFound
}
