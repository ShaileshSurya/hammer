package requesto

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

// ConstantErr
type constantErr string

// Error Constants
const (
	BuildError      = constantErr("Error While Building Request")
	RespDecodeError = constantErr("Error While Decoding Response to Into Struct")
)

// Error
const (
	RespDecodeErrorx = "Error While Decoding Response to Into Struct"
)
