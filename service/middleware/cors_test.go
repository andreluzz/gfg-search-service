package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCORSMethodGet(t *testing.T) {
	corsHandler := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	corsHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "should be equal")
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"), "should be equal")
	assert.Equal(t, "GET, POST, PATCH, PUT, DELETE, OPTIONS", w.Header().Get("Access-Control-Allow-Methods"), "should be equal")
	assert.Equal(t, "Accept, Content-Type, Content-Length, Accept-Encoding, X-Auth-Token, X-CSRF-Token, X-Service-Version, Authorization", w.Header().Get("Access-Control-Allow-Headers"), "should be equal")
}

func TestCORSMethodOptions(t *testing.T) {
	corsHandler := CORS(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodOptions, "/products", nil)
	corsHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "should be equal")
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"), "should be equal")
	assert.Equal(t, "GET, POST, PATCH, PUT, DELETE, OPTIONS", w.Header().Get("Access-Control-Allow-Methods"), "should be equal")
	assert.Equal(t, "Accept, Content-Type, Content-Length, Accept-Encoding, X-Auth-Token, X-CSRF-Token, X-Service-Version, Authorization", w.Header().Get("Access-Control-Allow-Headers"), "should be equal")
}
