package filedata

import (
	"errors"
	"mime"
	"net/http"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/rafaelmartins/filebin/internal/settings"
)

func (f *FileData) GetLexer() chroma.Lexer {
	if f.lexer != nil {
		return f.lexer
	}

	mt, _, err := mime.ParseMediaType(f.Mimetype)
	if err != nil {
		return nil
	}

	// exceptions
	if mt == "text/plain" {
		return lexers.Get("plaintext")
	}

	return lexers.MatchMimeType(mt)
}

func (f *FileData) ServeFileHighlight(w http.ResponseWriter, r *http.Request) error {
	s, err := settings.Get()
	if err != nil {
		return err
	}

	style := styles.Get(s.HighlightStyle)
	if style == nil {
		style = styles.Fallback
	}

	contents, err := f.readFile()
	if err != nil {
		return err
	}

	lexer := f.GetLexer()
	if lexer == nil {
		return errors.New("filedata: no lexer found")
	}

	iterator, err := lexer.Tokenise(nil, string(contents))
	if err != nil {
		return err
	}

	formatter := html.New(
		html.Standalone(true),
		html.WithLineNumbers(true),
		html.LinkableLineNumbers(true, "L"),
		html.LineNumbersInTable(true),
	)

	return formatter.Format(w, style, iterator)
}
