package auth

type IUser interface {
	GetID() string
	GetUserName() string
	GetFullName() string
}