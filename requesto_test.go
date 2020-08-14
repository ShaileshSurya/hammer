package requesto

import (
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

func TestGetHTTPClient(t *testing.T) {
	client := (&Requesto{HTTPClient: &http.Client{}}).getHTTPClient()
	if !reflect.DeepEqual(client, &http.Client{}) {
		t.Error("TestFailed: TestGetHTTPClient")
	}

	clientx := (&Requesto{}).getHTTPClient()
	if !reflect.DeepEqual(clientx, &http.Client{}) {
		t.Error("TestFailed: TestGetHTTPClient")
	}
}

func TestDebug(t *testing.T) {
	questo := (&Requesto{}).Debug()

	if !questo.debugMode {
		t.Error("Test Failed: TestDebug")
	}
}

func TestExecute(t *testing.T) {
	req := &Request{
		url:         "http://localhost:8081/",
		httpVerb:    POST,
		requestBody: []byte(`bodySample`),
	}

	_, err := (&Requesto{}).Execute(req)
	if err == nil {
		t.Error("Test Failed:TestExecute ")
	}

}
