package request

type BodyRequsetError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}