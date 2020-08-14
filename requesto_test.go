package requesto

import (
	"errors"
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
	// client := (&Requesto{HTTPClient: httpClient{}}).getHTTPClient()

	// val, ok := client.(httpClient)
	// if !ok {
	// 	t.Error("TestFailed: TestGetHTTPClient")
	// }

	// if !reflect.DeepEqual(val.client, &http.Client{}) {
	// 	t.Error("TestFailed: TestGetHTTPClient")
	// }

	clientx := (&Requesto{}).getHTTPClient()
	valx, _ := clientx.(httpClient)
	// fmt.Println(valx.client)
	// fmt.Println(&http.Client{})
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
	err error
}

func (m MockClient) Do(req *http.Request) (*http.Response, error) {
	if m.err != nil {
		return nil, m.err
	}
	return &http.Response{}, nil
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
