package upload

import (
	"fmt"
	"io"
	"mime"
	"net/http"
	"net/textproto"
	"path/filepath"

	"github.com/rafaelmartins/filebin/internal/highlight"
)

var (
	contentType = textproto.CanonicalMIMEHeaderKey("content-type")
)

func getMimeType(filename string, r io.ReadSeeker, header textproto.MIMEHeader) (string, error) {
	// try source types first
	m := highlight.GetMimeType(filename)
	if m != "" {
		return m, nil
	}

	// then try from OS mime database (at least on unix)
	m = mime.TypeByExtension(filepath.Ext(filename))
	if m != "" {
		return m, nil
	}

	// then try signature detection from net/http
	var buf [512]byte
	n, _ := io.ReadFull(r, buf[:])
	if _, err := r.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	m = http.DetectContentType(buf[:n])
	if m != "" {
		return m, nil
	}

	// then try mime type sent by http client
	for key, l := range header {
		if len(l) > 0 && contentType == key {
			return l[0], nil
		}
	}

	return "", fmt.Errorf("upload: failed to detect mime type: %s", filename)
}
