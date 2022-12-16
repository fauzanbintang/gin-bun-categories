package responses

// A R stand for standard response,
// it's intended to wrap every response data that will be returned to the client
type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
