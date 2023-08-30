package mysql

import (
    "github.com/22722679/douyin-project/model"
    "fmt"
)

func InserRegisterInfo(info model.RegisterInfo) (int64,error) {


  sqlStr := "insert into register_info(user_name, password) values(?, ?)"
  ret,err:=  db.Exec(sqlStr, info.UserName, info.Password)
  if err != nil{
      fmt.Printf("insert failed, err:%v\n", err)
       return 0,err;
  }

  ID,err:= ret.LastInsertId() // 新插入数据的id
  if err != nil {
      fmt.Printf("get lastinsert ID failed, err:%v\n", err)
      return 0,err
  }
  return ID,nil;
}
