package middleware

import (
	"log"
	"mime"
	"net/http"
)

func Logging(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("recv a %s request from %s", r.Method, r.RemoteAddr)
		handler.ServeHTTP(w, r)
	})
}

func Validating(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if mediaType != "application/json" {
			http.Error(w, "invalid Content-Type", http.StatusUnsupportedMediaType)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
