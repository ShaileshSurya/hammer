package hammer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Hammer ...
type Hammer struct {
	HTTPClient httpOperations
	// Finalize on LogFunc it should adhere to the users log level
	LogFunc   func(msg string)
	debugMode bool
}

// New ...
func New() *Hammer {
	req := &Hammer{}
	return req
}

// func (r *Hammer) logMessage(message string) {
// 	if r.debugMode {
// 		r.LogFunc(message)
// 	}
// }

// WithHTTPClient ...
func (r *Hammer) WithHTTPClient(hClient *http.Client) *Hammer {
	r.HTTPClient = httpClient{client: hClient}
	return r
}

// func (r *Hammer) clone() *Hammer {
// 	return &Hammer{
// 		HTTPClient: r.HTTPClient,
// 		LogFunc:    r.LogFunc,
// 		debugMode:  r.debugMode,
// 	}
// }

// Debug ...
func (r *Hammer) Debug() *Hammer {
	r.debugMode = true
	return r
}

// Execute ...
func (r *Hammer) Execute(req *Request) (*http.Response, error) {
	httpClient := r.getHTTPClient()
	return req.doRequest(httpClient)
}

// ExecuteInto ...
func (r *Hammer) ExecuteInto(req *Request, value interface{}) error {
	resp, err := r.Execute(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal([]byte(body), value)
	if err != nil {
		return err
	}
	return nil
}
