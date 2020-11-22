package highlight

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/styles"
	"github.com/rafaelmartins/filebin/internal/settings"
)

var (
	css []byte
)

func getFormatter() *html.Formatter {
	return html.New(
		html.WithClasses(true),
		html.WithLineNumbers(true),
		html.LinkableLineNumbers(true, "L"),
		html.LineNumbersInTable(true),
	)
}

func GenerateHTML(w io.Writer, r io.Reader, l chroma.Lexer) error {
	if l == nil {
		return ErrUnsupported
	}

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	iterator, err := chroma.Coalesce(l).Tokenise(nil, string(data))
	if err != nil {
		return err
	}

	return getFormatter().Format(w, styles.Fallback, iterator)
}

func GenerateCSS(w io.Writer) error {
	if css != nil {
		_, err := w.Write(css)
		return err
	}

	s, err := settings.Get()
	if err != nil {
		return err
	}

	style := styles.Get(s.HighlightStyle)
	if style == nil {
		style = styles.Fallback
	}

	var buf bytes.Buffer
	mw := io.MultiWriter(w, &buf)

	if err := getFormatter().WriteCSS(mw, style); err != nil {
		return err
	}

	css = buf.Bytes()

	return nil
}
