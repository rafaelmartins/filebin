package raw

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/rafaelmartins/filebin/internal/filedata"
)

type RawRenderer struct{}

func (h *RawRenderer) Supports(mimetype string) bool {
	return true
}

func (h *RawRenderer) Render(w http.ResponseWriter, r *http.Request, fd *filedata.FileData) error {
	disposition := "attachment"
	mt := strings.Split(fd.Mimetype, "/")
	if len(mt) > 0 && (mt[0] == "audio" || mt[0] == "image" || mt[0] == "video") {
		disposition = "inline"
	}

	w.Header().Set("Content-Type", fd.Mimetype)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`%s; filename="%s"`, disposition, fd.GetFilename()))
	w.Header().Set("X-Content-Type-Options", "nosniff")

	return fd.ServeData(w, r)
}
