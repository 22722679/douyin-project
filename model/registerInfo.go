package model

import "gorm.io/gorm"

type RegisterInfo struct{
  gorm.Model
  UserName       string     `json:"user_name,omitempty"`    //用户名
  Password       string     `json:"password,omitempty"`     //密码
}
