package dto

type LoginRes struct {
	Token     string `json:"token"`
	ExpiresIn uint64 `json:"expiresIn"`
}

type AuthRes struct {
	UserID       string `json:"userId"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    uint64 `json:"expiresIn"`
}

type UserRes struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	PhoneCode      string `json:"phoneCode"`
	PhoneNumber    string `json:"phoneNumber"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	NameAndSurname string `json:"nameAndSurname"`
}

type RegisterRes struct {
	PhoneRefCode string   `json:"phoneRefCode"`
	MailRefCode  string   `json:"mailRefCode"`
	Auth         *AuthRes `json:"auth"`
	UserInfo     *UserRes `json:"userInfo"`
}
