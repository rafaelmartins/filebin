package mime

import (
	"errors"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"path/filepath"

	"github.com/alecthomas/chroma/lexers"
	"github.com/rafaelmartins/filebin/internal/mime/magic"
)

var (
	contentType = textproto.CanonicalMIMEHeaderKey("content-type")
	errNotFound = errors.New("mime: type not found")
)

func detectFromData(r io.ReadSeeker) (string, error) {
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

	return "", errNotFound
}

func getChromaMimetype(filename string) string {
	lexer := lexers.Match(filename)
	if lexer == nil {
		return ""
	}
	mimes := lexer.Config().MimeTypes
	if len(mimes) == 0 {
		return ""
	}
	return mimes[0]
}

func detectFromFilename(filename string) (string, error) {
	// try source types first
	if m := getChromaMimetype(filename); m != "" {
		return m, nil
	}

	// then try from OS mime database (at least on unix)
	if m := mime.TypeByExtension(filepath.Ext(filename)); m != "" {
		return m, nil
	}

	return "", errNotFound
}

func Detect(f io.ReadSeeker, fh *multipart.FileHeader) (string, error) {
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
