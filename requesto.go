package requesto

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

// Execute ...
func (r *Requesto) Execute(req *Request) (*http.Response, error) {
	httpClient := r.getHTTPClient()

	switch req.httpVerb {
	case GET:
		request, err := http.NewRequest(GET, req.url, nil)
		if err != nil {
			return nil, errors.New("Log While creating request")
		}

		for k, v := range req.headers {
			request.Header.Set(k, v)
		}

		resp, err := httpClient.Do(request)
		if err != nil {
			return resp, err
		}
		return resp, err
	default:
		return nil, nil
	}
	return nil, nil
}

// ExecuteInto ...
func (r *Requesto) ExecuteInto(req *Request, value interface{}) error {
	resp, err := r.Execute(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), value)
	if err != nil {
		fmt.Println(err)
		return errors.New(RespDecodeErrorx)
	}
	return nil

}
