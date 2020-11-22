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
	w.Header().Set("Content-Type", fd.Mimetype)
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if !(strings.HasPrefix(fd.Mimetype, "audio/") || strings.HasPrefix(fd.Mimetype, "image/") || strings.HasPrefix(fd.Mimetype, "video/")) {
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fd.GetFilename()))
	}

	return fd.ServeData(w, r)
}
