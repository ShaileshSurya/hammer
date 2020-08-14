package requesto

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// Request ...
type Request struct {
	httpVerb string
	baseURL  string
	endpoint string
	url      string

	// Check if the Errors can be used as a array of errors.
	err               error
	headers           map[string]string
	requesto          *Requesto
	into              interface{}
	requestParams     map[string]string
	Requesto          *Requesto
	requestBodyParams map[string]interface{}
	requestBody       []byte
}

func initRequest() *Request {
	return &Request{
		headers:           make(map[string]string),
		requestParams:     make(map[string]string),
		requestBodyParams: make(map[string]interface{}),
	}
}

func (r *Request) checkError() error {
	if r.err != nil {
		return r.err
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

// WithRequestParams ...
func (r *Request) WithRequestParams(key string, value string) *Request {
	r.requestParams[key] = value
	return r
}

// WithRequestBodyParams ...
func (r *Request) WithRequestBodyParams(key string, value interface{}) *Request {
	r.requestBodyParams[key] = value
	return r
}

// WithURL ...
func (r *Request) WithURL(value string) *Request {
	r.url = value
	return r
}

// WithID ...
func (r *Request) WithID(ID string) *Request {
	switch r.httpVerb {
	case GET, PUT, DELETE:
		r.WithURL(r.url + "/" + ID)
		return r
	default:
		errMessage := fmt.Sprintf("Can not use WithID(string) for %s httpVerb", r.httpVerb)
		r.err = errors.New(errMessage)
		return r
	}
}

// Get ...
func (r *Request) Get() *Request {
	r.httpVerb = GET
	return r
}

// Head ...
func (r *Request) Head() *Request {
	r.httpVerb = HEAD
	return r
}

// Post ...
func (r *Request) Post() *Request {
	r.httpVerb = POST
	return r
}

// Put ...
func (r *Request) Put() *Request {
	r.httpVerb = PUT
	return r
}

// Patch ...
func (r *Request) Patch() *Request {
	r.httpVerb = PATCH
	return r
}

// Delete ...
func (r *Request) Delete() *Request {
	r.httpVerb = DELETE
	return r
}

// Connect ...
func (r *Request) Connect() *Request {
	r.httpVerb = CONNECT
	return r
}

// Options ...
func (r *Request) Options() *Request {
	r.httpVerb = OPTIONS
	return r
}

// Trace ...
func (r *Request) Trace() *Request {
	r.httpVerb = TRACE
	return r
}

// WithRequestBody ....
func (r *Request) WithRequestBody(body interface{}) *Request {
	bodyData, _ := json.Marshal(body)
	r.requestBody = bodyData
	return r
}

// Build will do basic validations.
func (r *Request) Build() (*Request, error) {
	if err := r.checkError(); err != nil {
		return nil, err
	}
	if "" == r.httpVerb {
		return nil, errors.New("No HttpVerb Provided")
	}
	if "" == r.url {
		return nil, errors.New("No Url Provided")
	}

	// Check if this can be moved to WithHeaders method
	if r.requestParams != nil {
		reqParamString := "?"
		for k, v := range r.requestParams {
			if reqParamString != "?" {
				reqParamString += "&"
			}
			reqParamString = reqParamString + k + `=` + v
		}
		r.url = r.url + reqParamString
	}

	// TODO: Move to a Function
	if len(r.requestBody) == 0 && len(r.requestBodyParams) != 0 {
		r.WithRequestBody(r.requestBodyParams)
	}

	return r, nil
}

// TODO : UT for this.
func (r *Request) doRequest(httpClient *http.Client) (*http.Response, error) {

	reqBody := func(body []byte) *bytes.Buffer {
		if len(body) != 0 {
			return bytes.NewBuffer(body)
		}
		return nil
	}(r.requestBody)

	request, err := http.NewRequest(r.httpVerb, r.url, reqBody)
	if err != nil {
		return nil, errors.New("Log While creating request")
	}
	for k, v := range r.headers {
		request.Header.Set(k, v)
	}

	command, _ := GetCurlCommand(request)

	r.Requesto.logMessage(command.String())

	resp, doerr := httpClient.Do(request)
	if doerr != nil {
		return resp, doerr
	}
	return resp, err
}

// Execute ...
// func (r *Request) Execute() (interface{}, error) {

// 	if err := r.checkError(); err != nil {
// 		return nil, err
// 	}

// 	httpClient := r.requesto.getHTTPClient()

// 	switch r.httpVerb {
// 	case GET:
// 		resp, err := httpClient.Get(r.baseURL + r.endpoint)
// 		if err != nil {
// 			return resp, err
// 		}
// 		if r.into != nil {
// 			body, err := ioutil.ReadAll(resp.Body)
// 			if err != nil {
// 				return nil, err
// 			}
// 			err = json.Unmarshal([]byte(body), r.into)
// 			if err != nil {
// 				return nil, errors.New(RespDecodeErrorx)
// 			}
// 			return r.into, nil
// 		}

// 	case PUT:

// 	case POST:

// 	case DELETE:

// 	case PATCH:

// 	default:

// 	}
// 	return nil, nil
// }
