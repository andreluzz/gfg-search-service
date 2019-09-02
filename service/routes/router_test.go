package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andreluzz/gfg-search-service/service/contextkeys"
	"github.com/andreluzz/gfg-search-service/service/storage"
	"github.com/stretchr/testify/assert"
)

func TestNewRouter(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	req.Header.Add("X-Service-Version", "v2")
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.wDWyyGem9YgXDDbH3Un7YYcTB8IcN_BE4BMmS1tvlnE")

	r := New("elasticsearch-mock-url", mockSearchIndex)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductsVersion1(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	ctx := context.WithValue(req.Context(), contextkeys.KeyAPIVersion, "v1")

	GetProducts(w, req.WithContext(ctx), mockSearchIndex, "")

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductsVersion2(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products", nil)
	ctx := context.WithValue(req.Context(), contextkeys.KeyAPIVersion, "v2")

	GetProducts(w, req.WithContext(ctx), mockSearchIndex, "")

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetProductsVersion2InvalidQueryString(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/products?page=aaaa", nil)
	ctx := context.WithValue(req.Context(), contextkeys.KeyAPIVersion, "v2")

	GetProducts(w, req.WithContext(ctx), mockSearchIndexInvalid, "")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetProductsInvalidMethod(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/products", nil)

	GetProducts(w, req, mockSearchIndex, "")

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func mockSearchIndexInvalid(esHost, index, q, filter, sort, page, limit string, client storage.HTTPClient) (*storage.Response, error) {
	return nil, fmt.Errorf("invalid query string parameters test")
}

func mockSearchIndex(esHost, index, q, filter, sort, page, limit string, client storage.HTTPClient) (*storage.Response, error) {
	mockData := `{
		"hits" : {
			"total" : 30,
			"hits" : [
				{
					"_id" : "14",
					"_source" : {
						"title" : "Suede Classic Plus Sneakers",
						"brand" : "Puma",
						"price" : 58.82,
						"stock" : 22
					}
				},
				{
					"_id" : "19",
					"_source" : {
						"title" : "3-Stripes Tee",
						"brand" : "Adidas",
						"price" : 20.73,
						"stock" : 6
					}
				}
			]
		}
	}`

	resp := &storage.Response{}
	if err := json.Unmarshal([]byte(mockData), resp); err != nil {
		return nil, err
	}

	return resp, nil
}
