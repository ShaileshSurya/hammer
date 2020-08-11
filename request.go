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
	err      constantErr
	headers  map[string]string
	requesto *Requesto
	into     interface{}
}

func (r *Request) checkError() error {
	if r.err != "" {
		return errors.New(string(r.err))
	}
	return nil
}

// WithHeaders ...
func (r *Request) WithHeaders(key string, value string) *Request {
	r.headers[key] = value
	return r
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
