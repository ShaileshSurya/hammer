package hammer

import (
	"net/http"
	"reflect"
)

type httpOperations interface {
	Do(*http.Request) (*http.Response, error)
}

type httpClient struct {
	client *http.Client
}

func (r *Hammer) getHTTPClient() httpOperations {

	if r.HTTPClient == nil || reflect.DeepEqual(r.HTTPClient, httpClient{}) {
		r.HTTPClient = httpClient{
			client: &http.Client{},
		}
	}
	return r.HTTPClient
}

func (h httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.client.Do(req)
}
