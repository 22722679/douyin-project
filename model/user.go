package model

//import "github.com/jinzhu/gorm"


//用户账号、密码信息
type User struct {
	//gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

//type UserI struct {
//	gorm.Model
//	Name string `json:name,omitempty`
//}