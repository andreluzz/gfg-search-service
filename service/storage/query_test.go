package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSearchBody(t *testing.T) {
	q := "product_title"
	filter := "brand:brand_name"
	sort := "price:desc,stock:asc"
	page := "2"
	limit := "20"

	generatedQuery, err := generateSearchBody(q, filter, sort, page, limit)

	expectedQuery := `{
	"query": {
		"bool": {
			"must": [
				{
					"query_string": {
						"query": "product_title"
					}
				},
				{
					"match": {
						"brand": "brand_name"
					}
				}
			]
		}
	},
	"sort": [
		{
			"price": {
				"order": "desc"
			}
		},
		{
			"stock": {
				"order": "asc"
			}
		}
	],
	"from": 20,
	"size": 20
}`

	assert.Equal(t, expectedQuery, string(generatedQuery), "should be equals")
	assert.NoError(t, err, "should not return error")
}

func TestGenerateSearchBodyInvalidPage(t *testing.T) {
	q := "product_title"
	filter := "brand:brand_name"
	sort := "price:desc,stock:asc"
	page := "invalid-number"
	limit := "20"

	query, err := generateSearchBody(q, filter, sort, page, limit)

	assert.Error(t, err, "should return an error message")
	assert.Nil(t, query, "Should return a nil query")
}

func TestGenerateSearchBodyInvalidLimit(t *testing.T) {
	q := "product_title"
	filter := "brand:brand_name"
	sort := "price:desc,stock:asc"
	page := "1"
	limit := "invalid-number"

	query, err := generateSearchBody(q, filter, sort, page, limit)

	assert.Error(t, err, "should return an error message")
	assert.Nil(t, query, "Should return a nil query")
}

type test struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name"`
	Value int    `json:"value"`
}

func TestParseSortTextFields(t *testing.T) {
	sort := "value:desc,name:desc"

	parsedSort := ParseSortTextFields(sort, &test{})

	assert.Equal(t, "value:desc,name.keyword:desc", parsedSort, "Should be equals")
}
