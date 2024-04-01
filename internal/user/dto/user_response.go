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
	ID          string `json:"id"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	FullName    string `json:"fullName"`
}

type RegisterRes struct {
	Auth     *AuthRes `json:"auth"`
	UserInfo *UserRes `json:"userInfo"`
}
