package filedata

import (
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

func (f *FileData) GenerateHTML(w http.ResponseWriter) error {
	fn, err := f.getFilenameData()
	if err != nil {
		return err
	}

	fp, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer fp.Close()

	return highlight.GenerateHTML(w, fp, f.GetLexer())
}
