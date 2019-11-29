package middlewares

import (
	"net/http"
)

// EnsureRequestIsJSON checks HTTP request for a Content-Type header of JSON and returns http.StatusBadRequest if not
func EnsureRequestIsJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if "application/json" == r.Header.Get("Content-Type") {
			next.ServeHTTP(w, r)
			return
		}
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	})
}
