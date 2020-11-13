package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rafaelmartins/filebin/internal/download"
	"github.com/rafaelmartins/filebin/internal/index"
	"github.com/rafaelmartins/filebin/internal/upload"
)

func usage(str string) {
	fmt.Fprintln(os.Stderr, "usage: filebin")
	if str != "" {
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "error:", str)
		os.Exit(1)
	}
	os.Exit(0)
}

func usageErr(err error) {
	str := ""
	if err != nil {
		str = err.Error()
	}
	usage(str)
}

func main() {
	dir, found := os.LookupEnv("FILEBIN_STORAGE_DIR")
	if !found {
		dir = "data"
	}
	st, err := os.Stat(dir)
	if err != nil {
		if !os.IsNotExist(err) {
			usageErr(err)
		}
		if err := os.MkdirAll(dir, 0777); err != nil {
			usageErr(err)
		}
	} else if !st.IsDir() {
		usage(fmt.Sprintf("%q is not a directory", dir))
	}

	d := download.Downloader{
		Dir:   dir,
		Style: "github",
	}
	if style, ok := os.LookupEnv("FILEBIN_HIGHLIGHT_STYLE"); ok {
		if style == "" {
			usage("FILEBIN_HIGHLIGHT_STYLE must not be empty")
		}
		d.Style = style
	}

	user := os.Getenv("FILEBIN_AUTH_USERNAME")
	if user == "" {
		usage("FILEBIN_AUTH_USERNAME is required")
	}
	pass := os.Getenv("FILEBIN_AUTH_PASSWORD")
	if pass == "" {
		usage("FILEBIN_AUTH_PASSWORD is required")
	}

	u := upload.Uploader{
		Dir:      dir,
		MaxSize:  10 * 1024 * 1024,
		IdLength: 8,
		BaseURL:  os.Getenv("FILEBIN_BASE_URL"),
		Realm:    "filebin",
		Username: user,
		Password: pass,
	}
	if idLength, ok := os.LookupEnv("FILEBIN_ID_LENGTH"); ok {
		i, err := strconv.ParseUint(idLength, 10, 32)
		if err != nil {
			usageErr(err)
		}
		if i < 8 {
			usage("FILEBIN_ID_LENGTH must be >= 8")
		}
		u.IdLength = uint(i)
	}
	if maxSize, ok := os.LookupEnv("FILEBIN_UPLOAD_MAX_SIZE_MB"); ok {
		m, err := strconv.ParseInt(maxSize, 10, 64)
		if err != nil {
			usageErr(err)
		}
		if m == 0 {
			usage("FILEBIN_UPLOAD_MAX_SIZE_MB must be > 0")
		}
		u.MaxSize = m * 1024 * 1024
	}
	if realm, ok := os.LookupEnv("FILEBIN_AUTH_REALM"); ok {
		if realm == "" {
			usage("FILEBIN_AUTH_REALM must not be empty")
		}
		u.Realm = realm
	}

	r := mux.NewRouter()
	r.HandleFunc("/", u.Upload).Methods("POST")
	r.HandleFunc("/", index.Index)
	r.HandleFunc("/{id}.txt", d.DownloadText)
	r.HandleFunc("/{id}", d.Download)

	listenAddr := ":8000"
	if la, ok := os.LookupEnv("FILEBIN_LISTEN_ADDR"); ok {
		listenAddr = la
	}

	fmt.Fprintf(os.Stderr, " * Listening on %s\n", listenAddr)
	if err := http.ListenAndServe(listenAddr, handlers.LoggingHandler(os.Stderr, r)); err != nil {
		usageErr(err)
	}
}
