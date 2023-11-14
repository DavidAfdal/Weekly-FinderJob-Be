package request

type CreateJobInput struct {
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description,omitempty" validate:"required"`
	Company      string `json:"company,omitempty" validate:"required"`
	ImageCompany string `json:"image,omitempty" validate:"required"`
	Category     string `json:"category,omitempty" validate:"required"`
	Status       string `json:"status,omitempty" validate:"required"`
	Location     string `json:"location,omitempty" validate:"required"`
	Salary       int64  `json:"salary,omitempty" validate:"required"`
	UserId       string `json:"userId,omitempty" validate:"required"`
}

type UpdateJobInput struct {
	Id           int    `json:"id"`
	Title        string `json:"title" validate:"required"`
	Description  string `json:"description,omitempty" validate:"required"`
	Company      string `json:"company,omitempty" validate:"required"`
	ImageCompany string `json:"image,omitempty" validate:"required"`
	Category     string `json:"category,omitempty" validate:"required"`
	Status       string `json:"status,omitempty" validate:"required"`
	Location     string `json:"location,omitempty" validate:"required"`
	Salary       int64  `json:"salary,omitempty" validate:"required"`
}