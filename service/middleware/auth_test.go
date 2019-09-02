package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthValidToken(t *testing.T) {
	authHandler := Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.wDWyyGem9YgXDDbH3Un7YYcTB8IcN_BE4BMmS1tvlnE")
	authHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "should be equal")
}

func TestAuthHeaderTokenNotFound(t *testing.T) {
	authHandler := Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	authHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "should be equal")
}

func TestAuthHeaderInvalidToken(t *testing.T) {
	authHandler := Auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Add("Authorization", "invalid-token")
	authHandler.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code, "should be equal")
}
