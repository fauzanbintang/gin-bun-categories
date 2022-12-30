package responses

type ErrorResponse struct {
	Status int   `json:"status"`
	Error  error `json:"error"`
}
