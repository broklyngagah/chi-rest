package request

// RegisterRequest ...
type RegisterRequest struct {
	Fname       string `json:"fname" validate:"required"`
	Lname       string `json:"lname" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Phone       string `json:"phone" validate:"required"`
	Password    string `json:"password" validate:"required"`
	RecheckPass string `json:"recheckPassword" validate:"required"`
}
