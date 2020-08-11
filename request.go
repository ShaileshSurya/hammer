package requesto

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// Request ...
type Request struct {
	httpVerb string
	baseURL  string
	endpoint string
	url      string
	err      constantErr
	headers  map[string]string
	requesto *Requesto
	into     interface{}
}

func initRequest() *Request {
	return &Request{
		headers: make(map[string]string),
	}
}

func (r *Request) checkError() error {
	if r.err != "" {
		return errors.New(string(r.err))
	}
	return nil
}

// RequestBuilder ...
func RequestBuilder() *Request {
	return initRequest()
}

// WithHeaders ...
func (r *Request) WithHeaders(key string, value string) *Request {
	r.headers[key] = value
	return r
}

// WithURL ...
func (r *Request) WithURL(value string) *Request {
	r.url = value
	return r
}

// Get ...
func (r *Request) Get() *Request {
	r.httpVerb = GET
	return r
}

// Build will do basic validations.
func (r *Request) Build() (*Request, error) {
	if "" == r.httpVerb {
		return nil, errors.New("No HttpVerb Provided")
	}
	if "" == r.url {
		return nil, errors.New("No Url Provided")
	}
	return r, nil
}

// Into ...
func (r *Request) Into(value interface{}) *Request {
	r.into = value
	return r
}

// Execute ...
func (r *Request) Execute() (interface{}, error) {

	if err := r.checkError(); err != nil {
		return nil, err
	}

	httpClient := r.requesto.getHTTPClient()

	switch r.httpVerb {
	case GET:
		resp, err := httpClient.Get(r.baseURL + r.endpoint)
		if err != nil {
			return resp, err
		}
		if r.into != nil {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal([]byte(body), r.into)
			if err != nil {
				return nil, errors.New(RespDecodeErrorx)
			}
			return r.into, nil
		}

	case PUT:

	case POST:

	case DELETE:

	case PATCH:

	default:

	}
	return nil, nil
}
