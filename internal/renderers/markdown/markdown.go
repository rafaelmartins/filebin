package markdown

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rafaelmartins/filebin/internal/filedata"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

type MarkdownRenderer struct{}

func (h *MarkdownRenderer) Supports(mimetype string) bool {
	return mimetype == "text/x-markdown"
}

func (h *MarkdownRenderer) Render(w http.ResponseWriter, r *http.Request, fd *filedata.FileData) error {
	f, err := fd.Read()
	if err != nil {
		return err
	}
	defer f.Close()

	src, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithUnsafe(),
		),
	)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	_, err = io.WriteString(w, `<!DOCTYPE html>
<html>
<head>
<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
</head>
<body>
`)
	if err != nil {
		return err
	}

	if err := md.Convert(src, w); err != nil {
		return err
	}

	_, err = io.WriteString(w, `</body>
</html>`)
	return err
}
