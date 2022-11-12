package model

type UserModel struct {
	BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	CommonTimestampsField
}

func (UserModel) TableName() string {
	return "users"
}
