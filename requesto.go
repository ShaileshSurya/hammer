package requesto

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Requesto ...
type Requesto struct {
	HTTPClient httpOperations
	// Finalize on LogFunc it should adhere to the users log level
	LogFunc   func(msg string)
	debugMode bool
}

// New ...
func New() *Requesto {
	req := &Requesto{}
	return req
}

// func (r *Requesto) logMessage(message string) {
// 	if r.debugMode {
// 		r.LogFunc(message)
// 	}
// }

// WithHTTPClient ...
func (r *Requesto) WithHTTPClient(hClient *http.Client) *Requesto {
	r.HTTPClient = httpClient{client: hClient}
	return r
}

// func (r *Requesto) clone() *Requesto {
// 	return &Requesto{
// 		HTTPClient: r.HTTPClient,
// 		LogFunc:    r.LogFunc,
// 		debugMode:  r.debugMode,
// 	}
// }

// Debug ...
func (r *Requesto) Debug() *Requesto {
	r.debugMode = true
	return r
}

// Execute ...
func (r *Requesto) Execute(req *Request) (*http.Response, error) {
	httpClient := r.getHTTPClient()
	req.Requesto = r
	return req.doRequest(httpClient)
}

// ExecuteInto ...
func (r *Requesto) ExecuteInto(req *Request, value interface{}) error {
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
