package filedata

import (
	"bytes"
	"io"
	"net/http"
	"os"

	"github.com/rafaelmartins/filebin/internal/highlight"
)

func (f *FileData) GetLexer() string {
	if f.lexer != nil {
		return *f.lexer
	}
	l := highlight.GetLexer(f.Mimetype)
	f.lexer = &l
	return l
}

func (f *FileData) ServeHTML(w http.ResponseWriter, r *http.Request) error {
	if f.html != nil {
		_, err := w.Write(f.html)
		return err
	}

	fn, err := f.getFilenameData()
	if err != nil {
		return err
	}

	fp, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer fp.Close()

	var buf bytes.Buffer
	mw := io.MultiWriter(w, &buf)

	if err := highlight.GenerateHTML(mw, fp, f.GetLexer()); err != nil {
		return err
	}

	f.html = buf.Bytes()

	return err
}
