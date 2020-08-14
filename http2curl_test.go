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

	funcCheckIfContains := func(superString, substring string) bool {
		return strings.Contains(strings.ToUpper(superString), strings.ToUpper(substring))
	}
	if err != nil {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	request.Header.Set("key", "value")
	request.SetBasicAuth("username", "password")

	com, err := GetCurlCommand(request)
	if err != nil {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	if !funcCheckIfContains(com.String(), GET) {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	if !funcCheckIfContains(com.String(), testURL) {
		t.Error("Test Failed: TestGetCurlCommand")
	}
	if !funcCheckIfContains(com.String(), "value") {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	if !funcCheckIfContains(com.String(), "key") {
		t.Error("Test Failed: TestGetCurlCommand")
	}

	if !funcCheckIfContains(com.String(), "Authorization") {
		t.Error("Test Failed: TestGetCurlCommand")
	}
}

func TestClose(t *testing.T) {
	err := nopCloser{}.Close()
	if err != nil {
		t.Error("Failed TestingClose")
	}
}
