package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/rafaelmartins/filebin/internal/basicauth"
	"github.com/rafaelmartins/filebin/internal/filedata"
	"github.com/rafaelmartins/filebin/internal/settings"
	"github.com/rafaelmartins/filebin/internal/utils"
	"github.com/rafaelmartins/filebin/internal/version"
)

var (
	logo = `  __ _ _      _     _
 / _(_) | ___| |__ (_)_ __
| |_| | |/ _ \ '_ \| | '_ \
|  _| | |  __/ |_) | | | | |
|_| |_|_|\___|_.__/|_|_| |_|
`
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	fmt.Fprintf(w, "%s\n", logo)
	fmt.Fprintf(w, "Version %s, running at %s\n\n", version.Version, r.Host)
	fmt.Fprintf(w, "Source code: https://github.com/rafaelmartins/filebin\n")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	s, err := settings.Get()
	if err != nil {
		utils.Error(w, err)
		return
	}

	// authentication
	if !basicauth.BasicAuth(w, r, s.AuthRealm, s.AuthUsername, s.AuthPassword) {
		return
	}

	fd, err := filedata.NewFromRequest(r)
	if err != nil {
		utils.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if s.BaseUrl != "" {
		fmt.Fprintf(w, "%s/%s\n", s.BaseUrl, fd.GetId())
	} else {
		fmt.Fprintf(w, "%s\n", fd.GetId())
	}
}

func getFile(w http.ResponseWriter, r *http.Request) *filedata.FileData {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.NotFound(w, r)
		return nil
	}

	fd, err := filedata.NewFromId(id)
	if err != nil {
		if err == filedata.ErrNotFound {
			http.NotFound(w, r)
			return nil
		}
		utils.Error(w, err)
		return nil
	}
	return fd
}

func File(w http.ResponseWriter, r *http.Request) {
	fd := getFile(w, r)
	if fd == nil {
		return
	}

	if fd.GetLexer() != "" {
		if err := highlightFile(w, fd); err != nil {
			utils.Error(w, err)
		}
		return
	}

	w.Header().Set("Content-Type", fd.Mimetype)
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if !(strings.HasPrefix(fd.Mimetype, "audio/") || strings.HasPrefix(fd.Mimetype, "image/") || strings.HasPrefix(fd.Mimetype, "video/")) {
		w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fd.GetFilename()))
	}

	if err := fd.ServeData(w, r); err != nil {
		utils.Error(w, err)
	}
}

func FileText(w http.ResponseWriter, r *http.Request) {
	fd := getFile(w, r)
	if fd == nil {
		return
	}

	if fd.GetLexer() == "" {
		utils.ErrorBadRequest(w)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if err := fd.ServeData(w, r); err != nil {
		utils.Error(w, err)
	}
}

func FileDownload(w http.ResponseWriter, r *http.Request) {
	fd := getFile(w, r)
	if fd == nil {
		return
	}

	w.Header().Set("Content-Type", fd.Mimetype)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, fd.GetFilename()))
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if err := fd.ServeData(w, r); err != nil {
		utils.Error(w, err)
	}
}
