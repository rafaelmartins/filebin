package download

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelmartins/filebin/internal/highlight"
	"github.com/rafaelmartins/filebin/internal/models"
	"github.com/rafaelmartins/filebin/internal/utils"
)

type Downloader struct {
	Dir   string
	Style string
}

func (d *Downloader) base(w http.ResponseWriter, r *http.Request) (string, *models.FileData, bool) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.NotFound(w, r)
		return "", nil, false
	}

	fdm := models.FileDataModel{Dir: d.Dir}

	fd := &models.FileData{}
	if err := fdm.Read(id, fd); err != nil {
		if err == models.ErrNotFound {
			http.NotFound(w, r)
			return "", nil, false
		}
		utils.Error(w, err)
		return "", nil, false
	}

	fn := fdm.GetFilename(id)

	return fn, fd, true
}

func (d *Downloader) Download(w http.ResponseWriter, r *http.Request) {
	fn, fd, ok := d.base(w, r)
	if !ok {
		return
	}

	// try syntax highlight first
	err := highlight.Highlight(w, r, d.Style, fn, fd)
	switch err {
	case nil:
		return
	case highlight.ErrNotFound:
		http.NotFound(w, r)
		return
	case highlight.ErrNoLexer:
	default:
		utils.Error(w, err)
		return
	}

	// serve raw file
	if fd.Mimetype != "" {
		w.Header().Set("Content-Type", fd.Mimetype)
		w.Header().Set("X-Content-Type-Options", "nosniff")
	}
	http.ServeFile(w, r, fn)
}

func (d *Downloader) DownloadText(w http.ResponseWriter, r *http.Request) {
	fn, fd, ok := d.base(w, r)
	if !ok {
		return
	}

	if !highlight.Validate(fd.Mimetype) {
		utils.ErrorBadRequest(w)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	http.ServeFile(w, r, fn)
}
