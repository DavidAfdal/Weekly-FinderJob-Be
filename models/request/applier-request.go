package request

type ApplierRequest struct {
	UserId string `json:"user_id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	JobId  int    `json:"job_id" validate:"required"`
}