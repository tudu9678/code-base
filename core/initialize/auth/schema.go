package auth

type Auth struct {
	UserId               string
	AccessToken          string
	RefreshToken         string
	AccessTokenDuration  int64
	RefreshTokenDuration int64
}

type TokenInfo struct {
	Iss       string
	Sub       string
	UserName  string
	Type      string
	Exp       int64
	CreatedAt int64
	SubID     string
}

type User struct {
	ID       string
	UserName string
	FullName string
}

func (u *User) GetID() string {
	return u.ID
}

func (u *User) GetUserName() string {
	return u.UserName
}

func (u *User) GetFullName() string {
	return u.FullName
}
