package mysql

//import "os/user"

import (
	"database/sql"

	//"douyin-project/model"
  "strconv"
  "fmt"
	"github.com/22722679/douyin-project/model"
)


func CheckUserName(username string) uint{
  // var registerInfo model.RegisterInfo
  // 用户名  注册表？？？
  var Id string
  sqlStr := "SELECT id from users WHERE username = ?"
  if err := db.Get(&Id,sqlStr, username);err != nil {
    if err == sql.ErrNoRows {
      return 0;
    }
    fmt.Printf("check ::%v",err);
    return 0; 
    
  }
  fmt.Printf("%v\n", Id);
  intNum, _ := strconv.Atoi(Id)
   return uint(intNum)
  // return registerInfo
}


func SelectUserName(userId int64) (string, error) {

	var Name string

	sqlStr := "select name from user_infos where id = ? "

	if err := db.Select(&Name, sqlStr, userId); err != nil {

		if err == sql.ErrNoRows {

			Name = "zxh1238"

		}

	}


	return Name, nil

}


func InsertUserInfo(userInfo model.User) error{
    sqlStr := "insert into users(username,password) values(?,?)"
    if _,err := db.Exec(sqlStr,userInfo.Username, userInfo.Password); err != nil{
      fmt.Printf("%v\n",err)
      return err
    }
  return nil
}

