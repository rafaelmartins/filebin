package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rafaelmartins/filebin/internal/basicauth"
	"github.com/rafaelmartins/filebin/internal/filedata"
	"github.com/rafaelmartins/filebin/internal/highlight"
	"github.com/rafaelmartins/filebin/internal/renderers"
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
	footer := ""
	if s, err := settings.Get(); err == nil {
		footer = s.IndexFooter
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
<title>filebin - %s</title>
</head>
<body>
<pre>
%s
Version %s, running at %s

Source code: <a href="https://github.com/rafaelmartins/filebin">https://github.com/rafaelmartins/filebin</a>
%s</pre>
</body>
</html>
`, r.Host, logo, version.Version, r.Host, footer)
}

func Robots(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	fmt.Fprintln(w, "User-agent: *")
	fmt.Fprintln(w, "Disallow: /")
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// authentication
	if !basicauth.BasicAuth(w, r) {
		return
	}

	fds, err := filedata.NewFromRequest(r)
	if err != nil {
		if fds == nil {
			utils.Error(w, err)
			return
		}

		log.Printf("error: %s", err)

		// with at least one valid upload we won't return error
		found := false
		for _, fd := range fds {
			if fd != nil {
				found = true
				break
			}
		}
		if !found {
			utils.ErrorBadRequest(w)
			return
		}
	}

	baseUrl := ""
	if s, err := settings.Get(); err == nil {
		baseUrl = s.BaseUrl
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	for _, fd := range fds {
		if fd == nil {
			fmt.Fprintf(w, "failed\n")
			continue
		}
		if baseUrl != "" {
			fmt.Fprintf(w, "%s/%s\n", baseUrl, fd.GetId())
		} else {
			fmt.Fprintf(w, "%s\n", fd.GetId())
		}
	}
}

func List(w http.ResponseWriter, r *http.Request) {
	// authentication
	if !basicauth.BasicAuth(w, r) {
		return
	}

	baseUrl := ""
	if s, err := settings.Get(); err == nil {
		baseUrl = s.BaseUrl
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	filedata.ForEach(func(fd *filedata.FileData) {
		if baseUrl != "" {
			fmt.Fprintf(w, "%s: %s (%s) -> %s/%s\n", fd.Timestamp, fd.Filename, fd.Mimetype, baseUrl, fd.GetId())
		} else {
			fmt.Fprintf(w, "%s: %s (%s) -> %s\n", fd.Timestamp, fd.Filename, fd.Mimetype, fd.GetId())
		}
	})
}

func Delete(w http.ResponseWriter, r *http.Request) {
	// authentication
	if !basicauth.BasicAuth(w, r) {
		return
	}

	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.NotFound(w, r)
		return
	}

	if err := filedata.Delete(id); err != nil {
		if err == filedata.ErrNotFound {
			http.NotFound(w, r)
			return
		}
		utils.Error(w, err)
		return
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

	renderer, err := renderers.Lookup(fd.Mimetype)
	if err != nil {
		utils.Error(w, err)
	}

	if err := renderer.Render(w, r, fd); err != nil {
		utils.Error(w, err)
	}
}

func FileText(w http.ResponseWriter, r *http.Request) {
	fd := getFile(w, r)
	if fd == nil {
		return
	}

	lexer, err := highlight.GetLexer(fd.Mimetype)
	if err != nil || lexer == nil {
		utils.ErrorBadRequest(w)
		return
	}

	if err := fd.Serve(w, r, fd.GetFilename(), "text/plain; charset=utf-8", fd.Timestamp, false); err != nil {
		utils.Error(w, err)
	}
}

func FileDownload(w http.ResponseWriter, r *http.Request) {
	fd := getFile(w, r)
	if fd == nil {
		return
	}

	if err := fd.Serve(w, r, fd.GetFilename(), fd.Mimetype, fd.Timestamp, true); err != nil {
		utils.Error(w, err)
	}
}

func FileJSON(w http.ResponseWriter, r *http.Request) {
	fd := getFile(w, r)
	if fd == nil {
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(fd); err != nil {
		utils.Error(w, err)
	}
}
