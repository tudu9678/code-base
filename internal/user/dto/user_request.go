package dto

// import "myapp/pkg/pb"

type CreateUserReq struct {
	Email       string `json:"email" validate:"required,email"`
	PhoneCode   string `json:"phoneCode" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required,phone_number"`
	Password    string `json:"password" validate:"required,password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Status      string `json:"status"`
	CreatedBy   string `json:"-"`
	UpdatedBy   string `json:"-"`
	Token       string `json:"captchaVerifyToken"`
	IP          string `json:"-"`
}

type LoginReq struct {
	Username string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
