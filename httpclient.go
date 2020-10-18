package hammer

import (
	"context"
	"net/http"
)

type httpOperations interface {
	Do(*http.Request) (*http.Response, error)
}

type httpClient struct {
	client *http.Client
}

func (h httpClient) Do(req *http.Request) (*http.Response, error) {
	return h.client.Do(req)
}

func httpDo(ctx context.Context, client httpOperations, req *http.Request, f func(*http.Response, error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)

	go func() {
		c <- f(client.Do(req))
	}()

	select {
	case <-ctx.Done():
		<-c // Wait for f to return.
		return ctx.Err()
	case err := <-c:
		return err
	}
}
