package model



import "gorm.io/gorm"





//用户账号、密码信息

type User struct {

	gorm.Model
	Username string `json:"username"`

	Password string `json:"password"`

  Token string `json:"token"`
}


