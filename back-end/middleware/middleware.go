package middleware

import (
	"JUALiND/helper"
	"fmt"
	"net/http"
)

func Method(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			helper.ErrorResponseJSON(w, fmt.Errorf("Method is Not Allowed"), "Method is Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func ContentType(value string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", value)
		next.ServeHTTP(w, r)
	})
}
