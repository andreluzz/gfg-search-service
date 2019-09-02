package storage

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchIndex(t *testing.T) {
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
			}, nil
		},
	}

	resp, err := SearchIndex("elasticsearch-url", "products", "", "", "", "", "", client)

	assert.NotNil(t, resp, "should not be nil")
	assert.NoError(t, err, "should be nil")
}

func TestSearchIndexInvalidResponseBody(t *testing.T) {
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       ioutil.NopCloser(bytes.NewBufferString("{,,")),
			}, nil
		},
	}

	resp, err := SearchIndex("elasticsearch-url", "products", "", "", "", "", "", client)

	assert.Nil(t, resp, "should be nil")
	assert.Error(t, err, "should not be nil")
}

func TestSearchIndexUnavailableElasticsearchServer(t *testing.T) {
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return nil, fmt.Errorf("invalid url")
		},
	}

	resp, err := SearchIndex("elasticsearch-url", "products", "", "", "", "", "", client)

	assert.Nil(t, resp, "should be nil")
	assert.Error(t, err, "should not be nil")
}

func TestSearchIndexInvalidResponseCode(t *testing.T) {
	client := &MockClient{
		DoFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusNotFound,
				Body:       ioutil.NopCloser(bytes.NewBufferString("{}")),
			}, nil
		},
	}

	resp, err := SearchIndex("elasticsearch-url", "products", "", "", "", "", "", client)

	assert.Nil(t, resp, "should be nil")
	assert.Error(t, err, "should not be nil")
}

func TestSearchIndexInvalidSearchParameters(t *testing.T) {
	resp, err := SearchIndex("elasticsearch-url", "products", "", "", "", "aaa", "aaa", nil)

	assert.Nil(t, resp, "should be nil")
	assert.Error(t, err, "should not be nil")
}

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	// just in case you want default correct return value
	return &http.Response{}, nil
}

type MockClientInvalidURL struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClientInvalidURL) Do(req *http.Request) (*http.Response, error) {
	return nil, fmt.Errorf("invalid url")
}
