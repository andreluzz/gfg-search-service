package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreluzz/gfg-search-service/service/contextkeys"
)

// Versioning validates the version number in all requests
func Versioning(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		version := r.Header.Get("X-Service-Version")
		if version == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: header key X-Service-Version not defined")
			return
		}
		if version != "v1" && version != "v2" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: invalid header key X-Service-Version only v1 or v2 are available")
			return
		}
		ctx := context.WithValue(r.Context(), contextkeys.KeyAPIVersion, version)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
