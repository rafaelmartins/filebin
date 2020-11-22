package html

import (
	"net/http"

	"github.com/rafaelmartins/filebin/internal/filedata"
)

type HtmlRenderer struct{}

func (h *HtmlRenderer) Supports(mimetype string) bool {
	return mimetype == "application/xhtml+xml" || mimetype == "text/html"
}

func (h *HtmlRenderer) Render(w http.ResponseWriter, r *http.Request, fd *filedata.FileData) error {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	return fd.ServeData(w, r)
}
