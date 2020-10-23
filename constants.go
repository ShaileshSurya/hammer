package hammer

// Constants for HTTP verbs.
const (
	GET     = "GET"
	HEAD    = "HEAD"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	OPTIONS = "OPTIONS"
	TRACE   = "TRACE"
)

// Pseudo-constants headers
var headers = struct {
	contextType string
}{"Content-Type"}

// Error
const (
	RespDecodeErrorx = "Error while decoding response to into struct"
)
