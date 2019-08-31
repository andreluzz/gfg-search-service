package model

import (
	"encoding/json"

	"github.com/andreluzz/gfg-search-service/service/storage"
)

// Product defines the object to store the product data
type Product struct {
	ID    string  `json:"id,omitempty"`
	Title string  `json:"title"`
	Brand string  `json:"brand"`
	Price float32 `json:"price"`
	Stock int     `json:"stock"`
}

// Products groups a list of products
type Products struct {
	Total   *int        `json:"total,omitempty"`
	Results interface{} `json:"results"`
}

// GetProducts return a list of products
func (p *Products) GetProducts(esHost, query, filter, sort, page, limit string, storageSearchIndexFunc storage.Search) error {
	sort = storage.ParseSortTextFields(sort, &Product{})
	storageResponse, err := storageSearchIndexFunc(esHost, "products", query, filter, sort, page, limit)
	if err != nil {
		return err
	}

	results := []Product{}

	for _, hit := range storageResponse.Hits.Hits {
		product := Product{
			ID: hit.ID,
		}
		sourceJSONBytes, err := json.Marshal(hit.Source)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(sourceJSONBytes, &product); err != nil {
			return err
		}
		results = append(results, product)
	}
	p.Total = &storageResponse.Hits.Total
	p.Results = results

	return nil
}

// GetProductsES return a list of products with the elasticsearch format
func (p *Products) GetProductsES(esHost, query, filter, sort, page, limit string, storageSearchIndexFunc storage.Search) error {
	sort = storage.ParseSortTextFields(sort, &Product{})
	storageResponse, err := storageSearchIndexFunc(esHost, "products", query, filter, sort, page, limit)
	if err != nil {
		return err
	}

	p.Results = storageResponse

	return nil
}
