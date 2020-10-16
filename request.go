package hammer

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
)

// Request ...
type Request struct {
	httpVerb string
	baseURL  string
	endpoint string
	url      string

	// Check if the Errors can be used as a array of errors.
	err     error
	headers map[string]string
	// TODO: use url.Values{} here.
	requestParams     map[string]string
	formParams        url.Values
	requestBodyParams map[string]interface{}
	requestBody       []byte
	ctx               context.Context
	basicAuth
}

type basicAuth struct {
	username string
	password string
}

func initRequest() *Request {
	return &Request{
		headers:           make(map[string]string),
		requestParams:     make(map[string]string),
		requestBodyParams: make(map[string]interface{}),
		ctx:               context.Background(),
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

// WithRequestBody ....
func (r *Request) WithRequestBody(body interface{}) *Request {
	bodyData, _ := json.Marshal(body)
	r.requestBody = bodyData
	return r
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

// WithContext ...
func (r *Request) WithContext(ctx context.Context) *Request {
	r.ctx = ctx
	return r
}

// WithBasicAuth ...
func (r *Request) WithBasicAuth(user, pass string) *Request {
	r.basicAuth = basicAuth{
		username: user,
		password: pass,
	}
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

// WithTemplate ...
func (r *Request) WithTemplate(tempRequest *Request) *Request {
	return tempRequest
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

// Build will do basic validations.
func (r *Request) Build() (*Request, error) {
	if err := r.checkError(); err != nil {
		return nil, err
	}
	if "" == r.httpVerb {
		return nil, errors.New("No HttpVerb Provided")
	}
	if "" == r.url {
		return nil, errors.New("No URL Provided")
	}

	// Check if this can be moved to WithHeaders method or using url.Values{}
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
// TODO : Break this function.
func (r *Request) doRequestWithContext(client httpOperations) (*http.Response, error) {

	reqBody := func(body []byte) *bytes.Buffer {
		if len(body) != 0 {
			return bytes.NewBuffer(body)
		}
		return &bytes.Buffer{}
	}(r.requestBody)

	request, err := http.NewRequest(r.httpVerb, r.url, reqBody)
	if err != nil {
		return nil, errors.New("Error While Creating Request")
	}
	for k, v := range r.headers {
		request.Header.Set(k, v)
	}

	//Add Basic auth
	if !reflect.DeepEqual(r.basicAuth, basicAuth{}) {
		request.SetBasicAuth(r.basicAuth.username, r.basicAuth.password)
		r.basicAuth = basicAuth{}
	}

	// command, _ := GetCurlCommand(request)
	// r.Hammer.logMessage(command.String())

	request = request.WithContext(r.ctx)
	var response *http.Response

	doerr := client.httpDo(r.ctx, request, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		response = resp
		return nil
	})
	if doerr != nil {
		return response, doerr
	}

	return response, err
}

// TODO : UT for this.
// TODO : Break this function.
func (r *Request) doRequest(client httpOperations) (*http.Response, error) {

	reqBody := func(body []byte) *bytes.Buffer {
		if len(body) != 0 {
			return bytes.NewBuffer(body)
		}
		return &bytes.Buffer{}
	}(r.requestBody)

	request, err := http.NewRequest(r.httpVerb, r.url, reqBody)
	if err != nil {
		return nil, errors.New("Error While Creating Request")
	}
	for k, v := range r.headers {
		request.Header.Set(k, v)
	}

	//Add Basic auth
	if !reflect.DeepEqual(r.basicAuth, basicAuth{}) {
		request.SetBasicAuth(r.basicAuth.username, r.basicAuth.password)
		r.basicAuth = basicAuth{}
	}

	// command, _ := GetCurlCommand(request)
	// r.Hammer.logMessage(command.String())

	resp, doerr := client.Do(request)
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

// 	httpClient := r.hammer.getHTTPClient()

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
