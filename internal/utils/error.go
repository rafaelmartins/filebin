package utils

import (
	"log"
	"net/http"
)

func Error(w http.ResponseWriter, err error) {
	if err != nil {
		// let's not expose the actual error
		log.Printf("error: %s", err)
		if _, ok := err.(*http.ProtocolError); ok {
			ErrorBadRequest(w)
			return
		}
		ErrorInternalServerError(w)
	}
}

func ErrorBadRequest(w http.ResponseWriter) {
	http.Error(w, "400 bad request", http.StatusBadRequest)
}

func ErrorInternalServerError(w http.ResponseWriter) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}
