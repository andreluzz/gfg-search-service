package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreluzz/gfg-search-service/service/contextkeys"

	"github.com/stretchr/testify/assert"
)

func TestVersioning(t *testing.T) {
	versioningHandler := Versioning(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value(contextkeys.KeyAPIVersion) == "v2" {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Add("X-Service-Version", "v2")
	versioningHandler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code, "should be equal")
}

func TestVersioningHeaderNotDefined(t *testing.T) {
	versioningHandler := Versioning(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	versioningHandler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "should be equal")
}

func TestVersioningInvalidValue(t *testing.T) {
	versioningHandler := Versioning(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Add("X-Service-Version", "invalid-version")
	versioningHandler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code, "should be equal")
}
