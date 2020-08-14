package requesto

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Requesto ...
type Requesto struct {
	HTTPClient *http.Client
	// Finalize on LogFunc it should adhere to the users log level
	LogFunc   func(msg string)
	debugMode bool
}

// New ...
func New() *Requesto {
	req := &Requesto{}
	return req
}

func (r *Requesto) logMessage(message string) {
	if r.debugMode {
		r.LogFunc(message)
	}
}

// func (r *Requesto) clone() *Requesto {
// 	return &Requesto{
// 		HTTPClient: r.HTTPClient,
// 		LogFunc:    r.LogFunc,
// 		debugMode:  r.debugMode,
// 	}
// }

func (r *Requesto) getHTTPClient() *http.Client {
	if r.HTTPClient == nil {
		r.HTTPClient = &http.Client{}
	}
	return r.HTTPClient
}

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
		r.logMessage("Error while executing request")
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logMessage("Error while reading response body")
		return err
	}

	err = json.Unmarshal([]byte(body), value)
	if err != nil {
		r.logMessage("Error while decoding response body into struct")
		return err
	}
	return nil
}
