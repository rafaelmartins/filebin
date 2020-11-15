package highlight

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
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

func GenerateHTML(w io.Writer, r io.Reader, lexer string) error {
	l := lexers.Get(lexer)
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

	_, err = io.WriteString(w, `<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
<style type="text/css">
`)
	if err != nil {
		return err
	}

	if err := generateCSS(w); err != nil {
		return err
	}

	_, err = io.WriteString(w, `</style>
</head>
<body>
`)
	if err != nil {
		return err
	}

	if err := getFormatter().Format(w, styles.Fallback, iterator); err != nil {
		return err
	}

	_, err = io.WriteString(w, `</body>
</html>`)
	if err != nil {
		return err
	}

	return nil
}

func generateCSS(w io.Writer) error {
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
