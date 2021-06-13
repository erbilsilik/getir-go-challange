package middleware

import (
	"mime"
	"net/http"
)

func EnforceJSONMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	contentType := r.Header.Get("Content-Type")

	if contentType != "" {
		mt, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(rw, "Malformed Content-Type header", http.StatusBadRequest)
			return
		}

		if mt != "application/json" {
			http.Error(rw, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
			return
		}
	}

	next(rw, r)
}
