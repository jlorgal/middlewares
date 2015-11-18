package middlewares

import (
	"log"
	"net/http"
)

func LoggingHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
		log.Printf("%s %s", r.Method, r.URL)
	}
}
