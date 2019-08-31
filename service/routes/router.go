package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreluzz/gfg-search-service/service/contextkeys"
	"github.com/andreluzz/gfg-search-service/service/middleware"
	"github.com/andreluzz/gfg-search-service/service/model"
	"github.com/andreluzz/gfg-search-service/service/storage"
)

// New initializes the router and register the endpoints
func New(esHost string, storageFunc storage.Search) *http.ServeMux {
	router := http.NewServeMux()
	router.Handle("/products", middleware.Default(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetProducts(w, r, storageFunc, esHost)
	})))
	return router
}

// GetProducts routes the request and returns the response with a filtered list of products
func GetProducts(w http.ResponseWriter, r *http.Request, storageFunc storage.Search, esHost string) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	query := r.URL.Query().Get("q")
	filter := r.URL.Query().Get("filter")
	sort := r.URL.Query().Get("sort")
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	version := r.Context().Value(contextkeys.KeyAPIVersion)

	products := model.Products{}

	var err error
	switch version {
	case "v1":
		err = products.GetProductsES(esHost, query, filter, sort, page, limit, storageFunc)
	case "v2":
		err = products.GetProducts(esHost, query, filter, sort, page, limit, storageFunc)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
