package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

//Search defines a storage search interface to simplify the unit tests
type Search func(esHost, index, q, filter, sort, page, limit string, client HTTPClient) (*Response, error)

// HTTPClient interaface to simplify the unit tests
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Response defines the struct to parse the elasticsearch response
type Response struct {
	Hits struct {
		Total int   `json:"total"`
		Hits  []Hit `json:"hits"`
	} `json:"hits"`
}

// Hit defines a documento from the elasticsearch response
type Hit struct {
	ID     string      `json:"_id"`
	Source interface{} `json:"_source"`
}

// SearchIndex returns a list of documents from the defined index
// applying filter, query, sort and page
func SearchIndex(esHost, index, q, filter, sort, page, limit string, client HTTPClient) (*Response, error) {
	query, err := generateSearchBody(q, filter, sort, page, limit)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/_search?filter_path=hits.total,hits.hits._id,hits.hits._source", esHost, index)
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer(query))
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error getting data from elasticsearch: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid response code: %d", res.StatusCode)
	}

	resp := &Response{}
	if err := json.NewDecoder(res.Body).Decode(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
