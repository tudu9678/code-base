package dto

// import "myapp/pkg/pb"

type CreateUserReq struct {
	Email       string `json:"email" validate:"required,email"`
	UserName    string `json:"userName" `
	PhoneNumber string `json:"phoneNumber" validate:"required,phone_number"`
	Password    string `json:"password" validate:"required,password"`
	FullName    string `json:"fullName"`
	Dob         string `json:"dob"`
}

type LoginReq struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
