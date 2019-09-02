package model

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"

	"github.com/andreluzz/gfg-search-service/service/storage"
	"github.com/stretchr/testify/assert"
)

func TestGetProducts(t *testing.T) {
	products := Products{}
	err := products.GetProducts("elasticsearch-url", "", "", "", "", "", mockSearchIndex)

	assert.Equal(t, 2, *products.Total, "should be equals")
	assert.NoError(t, err, "should be nil")
}

func TestGetProductsInvalidSourceInterface(t *testing.T) {
	products := Products{}
	err := products.GetProducts("elasticsearch-url", "", "", "", "", "", mockSearchIndexInvalidSourceInterface)

	assert.Error(t, err, "should be nil")
}

func TestGetProductsInvalidSource(t *testing.T) {
	products := Products{}
	err := products.GetProducts("elasticsearch-url", "", "", "", "", "", mockSearchIndexInvalidSource)

	assert.Error(t, err, "should be nil")
}

func TestGetProductsInvalidESResponse(t *testing.T) {
	products := Products{}
	err := products.GetProducts("elasticsearch-url", "", "", "", "", "", mockSearchIndexInvalid)

	assert.Nil(t, products.Results, "should be nil")
	assert.Error(t, err, "should be nil")
}

func TestGetProductsES(t *testing.T) {
	products := Products{}
	err := products.GetProductsES("elasticsearch-url", "", "", "", "", "", mockSearchIndex)

	assert.Nil(t, products.Total, "should be nil")
	assert.NoError(t, err, "should be nil")
}

func TestGetProductsESInvalidESResponse(t *testing.T) {
	products := Products{}
	err := products.GetProductsES("elasticsearch-url", "", "", "", "", "", mockSearchIndexInvalid)

	assert.Nil(t, products.Results, "should be nil")
	assert.Error(t, err, "should be nil")
}

func mockSearchIndexInvalid(esHost, index, q, filter, sort, page, limit string, client storage.HTTPClient) (*storage.Response, error) {
	return nil, fmt.Errorf("invalid query string parameters test")
}

func mockSearchIndexInvalidSourceInterface(esHost, index, q, filter, sort, page, limit string, client storage.HTTPClient) (*storage.Response, error) {
	resp := &storage.Response{}
	h := storage.Hit{
		ID:     "0001",
		Source: math.Inf(1),
	}
	resp.Hits.Hits = append(resp.Hits.Hits, h)
	return resp, nil
}

func mockSearchIndexInvalidSource(esHost, index, q, filter, sort, page, limit string, client storage.HTTPClient) (*storage.Response, error) {
	mockData := `{
		"hits" : {
			"total" : 2,
			"hits" : [
				{
					"_id" : "14",
					"_source" : {
						"title" : "Suede Classic Plus Sneakers",
						"brand" : 1,
						"price" : 58.82,
						"stock" : 22
					}
				},
				{
					"_id" : "19",
					"_source" : {
						"title" : "3-Stripes Tee",
						"brand" : 2,
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

func mockSearchIndex(esHost, index, q, filter, sort, page, limit string, client storage.HTTPClient) (*storage.Response, error) {
	mockData := `{
		"hits" : {
			"total" : 2,
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
