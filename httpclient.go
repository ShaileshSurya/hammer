package hammer

import (
	"context"
	"net/http"
	"reflect"
)

type httpOperations interface {
	Do(*http.Request) (*http.Response, error)
	httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error
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

func (h httpClient) httpDo(ctx context.Context, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)

	go func() {
		c <- f(h.client.Do(req))
	}()

	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}

func (h httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.client.Do(req)
}
