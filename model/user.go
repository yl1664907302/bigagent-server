package model

type User struct {
	Username    string `gorm:"column:username" json:"username"`
	Email       string `gorm:"column:email" json:"email"`
	Password    string `gorm:"column:password" json:"password"`
	Role        string `gorm:"column:role" json:"role"`
	RoleId      string `gorm:"column:roleId" json:"roleId"`
	Permissions string `gorm:"column:permissions" json:"permissions"`
}

func (*User) TableName() string {
	return "user"
}
