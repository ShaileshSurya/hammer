package requesto

import (
	"bytes"
	"net/http"
	"strings"
	"testing"
)

const (
	testURL = "http://localhost:8081/"
)

func TestGetCurlCommand(t *testing.T) {

	request, err := http.NewRequest(GET, testURL, bytes.NewBuffer([]byte("sdf")))
	if err != nil {
		t.Error("Test Failed: TestGetCurlCommand")
	}
	request.Header.Set("key", "value")
	com, err := GetCurlCommand(request)
	if err != nil {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	if !strings.Contains(com.String(), GET) {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	if !strings.Contains(com.String(), testURL) {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	if !strings.Contains(com.String(), "value") {
		t.Error("Test Failed: TestGetCurlCommand")
	}
}
