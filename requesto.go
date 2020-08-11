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

/**

client := &http.Client{}
  request, err := http.NewRequest("GET", "http://example.com", nil)

          if err != nil {
                  log.Fatalln(err)
          }
  request.Header.Set("User-Agent", "[your user-agent name]")
  resp, err := client.Do(req)

*/
// Execute ...
func (r *Requesto) Execute(req *Request) (*http.Response, error) {
	httpClient := r.getHTTPClient()

	switch req.httpVerb {
	case GET:
		fmt.Println("~~~~~~~~~~~~~~~~HERE in GET Switch~~~~~~~~~~~~")
		request, err := http.NewRequest(GET, req.url, nil)
		if err != nil {
			fmt.Println("Log While creating request")
			return nil, errors.New("Log While creating request")
		}
		for k, v := range req.headers {
			request.Header.Set(k, v)
		}
		PrettyPrint(request)
		resp, err := httpClient.Do(request)
		if err != nil {
			fmt.Println("~~~~~~~~~~~~~~~~HERE in Error~~~~~~~~~~~~")
			return resp, err
		}
		fmt.Println("~~~~~~~~~~~~~~~~HERE After Error~~~~~~~~~~~~")
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
		return errors.New(RespDecodeErrorx)
	}
	return nil

}