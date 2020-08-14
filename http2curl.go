package requesto

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

// CurlCommand contains exec.Command compatible slice + helpers
type CurlCommand []string

// append appends a string to the CurlCommand
func (c *CurlCommand) append(newSlice ...string) {
	*c = append(*c, newSlice...)
}

// String returns a ready to copy/paste command
func (c *CurlCommand) String() string {
	return strings.Join(*c, " ")
}

// nopCloser is used to create a new io.ReadCloser for req.Body
type nopCloser struct {
	io.Reader
}

func bashEscape(str string) string {
	return `'` + strings.Replace(str, `'`, `'\''`, -1) + `'`
}

func (nopCloser) Close() error { return nil }

// GetCurlCommand returns a CurlCommand corresponding to an http.Request
func GetCurlCommand(req *http.Request) (*CurlCommand, error) {
	command := CurlCommand{}

	command.append("\n")
	command.append("curl")
	command.append("-X", bashEscape(req.Method))

	if req.Body != nil {
		body, _ := ioutil.ReadAll(req.Body)
		req.Body = nopCloser{bytes.NewBuffer(body)}
		if len(string(body)) > 0 {
			bodyEscaped := bashEscape(string(body))
			command.append("\n")
			command.append("-d", bodyEscaped)
		}
	}

	var keys []string

	for k := range req.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		command.append("\n")
		command.append("-H", bashEscape(fmt.Sprintf("%s: %s", k, strings.Join(req.Header[k], " "))))
	}

	command.append("\n")
	command.append(bashEscape(req.URL.String()))
	command.append("\n")
	command.append("\n")

	return &command, nil
}
