package raw

import (
	"net/http"
	"strings"

	"github.com/rafaelmartins/filebin/internal/filedata"
)

type RawRenderer struct{}

func (h *RawRenderer) Supports(mimetype string) bool {
	return true
}

func (h *RawRenderer) Render(w http.ResponseWriter, r *http.Request, fd *filedata.FileData) error {
	attachment := true
	mt := strings.Split(fd.Mimetype, "/")
	if len(mt) > 0 && (mt[0] == "audio" || mt[0] == "image" || mt[0] == "video") {
		attachment = false
	}

	return fd.ServeData(w, r, fd.Mimetype, fd.GetFilename(), attachment)
}
