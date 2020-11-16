package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rafaelmartins/filebin/internal/settings"
	"github.com/rafaelmartins/filebin/internal/views"
)

func usage(err error) {
	fmt.Fprintln(os.Stderr, "usage: filebin")
	if err != nil {
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "error:", err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

func main() {
	s, err := settings.Get()
	if err != nil {
		usage(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", views.Upload).Methods("POST")
	r.HandleFunc("/", views.Index)
	r.HandleFunc("/download/{id}", views.FileDownload)
	r.HandleFunc("/{id}.txt", views.FileText)
	r.HandleFunc("/{id}", views.File)

	fmt.Fprintf(os.Stderr, " * Listening on %s\n", s.ListenAddr)
	if err := http.ListenAndServe(s.ListenAddr, handlers.LoggingHandler(os.Stderr, r)); err != nil {
		usage(err)
	}
}
