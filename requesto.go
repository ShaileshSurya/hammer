package requesto

import "net/http"

// Requesto ...
type Requesto struct {
	BaseURL            string
	HTTPClient         *http.Client
	DefaultContentType string
	DefaultHeaders     map[string]string
}

// New ...
func New(req *Requesto) *Requesto {
	if req.DefaultHeaders == nil {
		req.DefaultHeaders = make(map[string]string)
	}
	return req
}

func (r *Requesto) clone() *Requesto {
	return &Requesto{
		BaseURL:            r.BaseURL,
		HTTPClient:         r.HTTPClient,
		DefaultContentType: r.DefaultContentType,
		DefaultHeaders:     r.DefaultHeaders,
	}
}

// Get ...
func (r *Requesto) Get(url string) (req *Request) {
	r = New(r)
	return &Request{
		baseURL:  r.BaseURL,
		endpoint: url,
		httpVerb: GET,
		headers:  r.DefaultHeaders,
		requesto: r.clone(),
	}
}

func (r *Requesto) getHTTPClient() *http.Client {
	if r.HTTPClient == nil {
		r.HTTPClient = &http.Client{}
	}
	return r.HTTPClient
}
