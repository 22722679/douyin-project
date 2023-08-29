package mysql



//import "os/user"



import (

	"database/sql"

	//"douyin-project/model"

	"fmt"

)



func SelectUserName(userId int64) (string, error) {

	var Name string

	sqlStr := fmt.Sprintf("select name from user_infos where id = %d", userId)

	if err := db.Select(&Name, sqlStr); err != nil {

		if err == sql.ErrNoRows {

			Name = "zxh1238"

		}

	}

	//db.Select("name").Where("id = ?",userId).First(&Name)

	return Name, nil

}

