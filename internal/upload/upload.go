package upload

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/rafaelmartins/filebin/internal/basicauth"
	"github.com/rafaelmartins/filebin/internal/id"
	"github.com/rafaelmartins/filebin/internal/models"
	"github.com/rafaelmartins/filebin/internal/utils"
)

type Uploader struct {
	Dir      string
	MaxSize  int64
	IdLength uint
	BaseURL  string
	Realm    string
	Username string
	Password string
}

func (u *Uploader) Upload(w http.ResponseWriter, r *http.Request) {
	// authentication
	if !basicauth.BasicAuth(w, r, u.Realm, u.Username, u.Password) {
		return
	}

	// we store all file data in memory, instead of letting the library use temp files.
	// this is intended to be used as a private service, it should not be an issue.
	if err := r.ParseMultipartForm(u.MaxSize); err != nil {
		utils.Error(w, err)
		return
	}

	f, h, err := r.FormFile("file")
	if err != nil {
		utils.Error(w, err)
		return
	}
	defer f.Close()

	if h.Size > u.MaxSize {
		utils.Error(w, fmt.Errorf("upload: file is too big (%s): %d", h.Filename, h.Size))
		return
	}

	mt, err := getMimeType(h.Filename, f, h.Header)
	if err != nil {
		utils.Error(w, err)
		return
	}

	d := &models.FileData{
		Filename:  h.Filename,
		Mimetype:  mt,
		Size:      h.Size,
		Timestamp: time.Now().UTC(),
	}

	fdm := models.FileDataModel{Dir: u.Dir}
	xid := ""
	for {
		var err error
		xid, err = id.Generate(u.IdLength)
		if err != nil {
			utils.Error(w, err)
			return
		}

		if err := fdm.Write(xid, d); err != nil {
			if err != models.ErrDuplicated {
				utils.Error(w, err)
				return
			}
		} else {
			break
		}
	}

	fn := fdm.GetFilename(xid)
	t, err := os.Create(fn)
	if err != nil {
		utils.Error(w, err)
		return
	}

	n, err := io.Copy(t, f)
	if err != nil {
		utils.Error(w, err)
		t.Close()
		os.RemoveAll(fn)
		return
	}
	if n != d.Size {
		utils.Error(w, fmt.Errorf("upload: failed to finish reading file (%s): %d", d.Filename, d.Size))
		t.Close()
		os.RemoveAll(fn)
		return
	}
	if err := t.Close(); err != nil {
		utils.Error(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	if u.BaseURL != "" {
		fmt.Fprintf(w, "%s/%s\n", u.BaseURL, xid)
	} else {
		fmt.Fprintf(w, "%s\n", xid)
	}
}
