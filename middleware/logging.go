package middleware

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Println(request.RequestURI)
		next.ServeHTTP(writer, request)
	})
}