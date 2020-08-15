package requesto

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	req := New()
	if !reflect.DeepEqual(req, &Requesto{}) {
		t.Error("Test Failed:TestNew")
	}
}

//func Implements(V Type, T *Interface) bool
func TestGetHTTPClient(t *testing.T) {
	client := (&Requesto{HTTPClient: httpClient{}}).getHTTPClient()

	val, ok := client.(httpClient)
	if !ok {
		t.Error("TestFailed: TestGetHTTPClient")
	}

	if !reflect.DeepEqual(val.client, &http.Client{}) {
		t.Error("TestFailed: TestGetHTTPClient")
	}

	clientx := (&Requesto{}).getHTTPClient()
	valx, _ := clientx.(httpClient)
	if !reflect.DeepEqual(valx.client, &http.Client{}) {
		t.Error("TestFailed: TestGetHTTPClient")
	}
}

func TestDebug(t *testing.T) {
	questo := (&Requesto{}).Debug()

	if !questo.debugMode {
		t.Error("Test Failed: TestDebug")
	}
}

type MockClient struct {
	err      error
	response *http.Response
}

func (m MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.response, nil
}

func TestExecute(t *testing.T) {
	req := &Request{
		url:         "http://localhost:8081/",
		httpVerb:    POST,
		requestBody: []byte(`bodySample`),
	}

	requesto := &Requesto{
		HTTPClient: MockClient{
			err: errors.New("Error"),
		},
	}
	_, err := requesto.Execute(req)
	if err == nil {
		t.Error("Test Failed:TestExecute ")
	}

	requesto = &Requesto{
		HTTPClient: MockClient{},
	}
	_, xerr := requesto.Execute(req)
	if xerr != nil {
		t.Error("Test Failed:TestExecute ")
	}
}

func TestExecuteInto(t *testing.T) {
	req := &Request{
		url:         "http://localhost:8081/",
		httpVerb:    POST,
		requestBody: []byte(`bodySample`),
	}

	body := []byte(`{"name":"name","job_title":"jobTitle1","job_title2":"jobTitle2","Nested":{"field":"filed1","field2":0,"field3":0}}`)

	requesto := &Requesto{
		HTTPClient: MockClient{
			//err: errors.New("Error"),
			response: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: struct {
					io.Reader
					io.Closer
				}{
					io.MultiReader(bytes.NewReader(body), http.NoBody),
					http.NoBody,
				},
			},
		},
	}
	var emp Employee
	err := requesto.ExecuteInto(req, &emp)
	if err != nil {
		t.Error("Test Failed:TestExecute ")
	}
}

func TestExecuteIntoErrExecute(t *testing.T) {
	req := &Request{
		url:         "http://localhost:8081/",
		httpVerb:    POST,
		requestBody: []byte(`bodySample`),
	}
	requesto := &Requesto{
		HTTPClient: MockClient{
			err: errors.New("Error"),
		},
	}
	var emp Employee
	err := requesto.ExecuteInto(req, &emp)
	if err == nil {
		t.Error("Test Failed:TestExecute ")
	}
}

func TestExecuteIntoErrMarshal(t *testing.T) {
	req := &Request{
		url:         "http://localhost:8081/",
		httpVerb:    POST,
		requestBody: []byte(`bodySample`),
	}
	body := []byte(`{"name":"name","job_title":"jobTitle1""job_title2":"jobTitle2","Nested":{"field":"filed1","field2":0,"field3":0`)

	requesto := &Requesto{
		HTTPClient: MockClient{
			//err: errors.New("Error"),
			response: &http.Response{
				Status:     "200 OK",
				StatusCode: 200,
				Body: struct {
					io.Reader
					io.Closer
				}{
					io.MultiReader(bytes.NewReader(body), http.NoBody),
					http.NoBody,
				},
			},
		},
	}
	var emp Employee
	err := requesto.ExecuteInto(req, &emp)
	if err == nil {
		t.Error("Test Failed:TestExecute ")
	}
}

func TestWithHTTPClient(t *testing.T) {
	client := New().WithHTTPClient(&http.Client{})

	createdClient := (client.HTTPClient.(httpClient)).client
	if !reflect.DeepEqual(createdClient, &http.Client{}) {
		t.Error("test Failed: TestWithHTTPClient")
	}
}
