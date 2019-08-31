package middleware

import (
	"net/http"
)

// Default use Auth and Versioning at the handler func
func Default(f http.HandlerFunc) http.HandlerFunc {
	return CORS(Logging(Auth(Versioning(f))))
}
