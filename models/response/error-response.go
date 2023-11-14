package response

type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Errors  interface{} `json:"errors"`
}
