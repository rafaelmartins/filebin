package markdown

import (
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
	f, err := fd.OpenData()
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
		),
	)

	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	return md.Convert(src, w)
}
