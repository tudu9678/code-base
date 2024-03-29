package model

// User ...
type User struct {
	BaseModel
	Email       string `gorm:"index;unique"`
	Password    string
	PhoneCode   string `gorm:"uniqueIndex:idx_users_phone"`
	PhoneNumber string `gorm:"uniqueIndex:idx_users_phone"`
	UnitHolder  string
	FirstName   string
	LastName    string
}

// TableName ...
func (User) TableName() string {
	return "users"
}
