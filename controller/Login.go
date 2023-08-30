package controller

import (

	//"database/sql"

	"github.com/22722679/douyin-project/model"
	"github.com/22722679/douyin-project/util"

	"github.com/22722679/douyin-project/mysql"

	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)



func Login(ctx *gin.Context) {
  
	db := mysql.GETDB()

	var requestUser model.User

	ctx.Bind(&requestUser)
  
  requestUser.Username,_ = ctx.GetQuery("username");
  requestUser.Password,_ = ctx.GetQuery("password");
  
	Name := requestUser.Username

	password := requestUser.Password

	//数据检验

	if len(Name) < 6 {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{

			"status_code": 421,

			"status_msg": "账户必须为大于6位",

			"user_id": 0,

			"token": "user",
		})

		return

	}

	if len(password) < 6 {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{

			"status_code": 420,

			"status_msg": "密码不能少于6位",

			"user_id": 0,

			"token": "password don`t small 6",
		})

		return

	}

	//判断账户是否存在

	var user model.User

	sqlStr := "select name from users"

	if err := db.Select(&user, sqlStr); err == nil {

		//if err == sql.ErrNoRows {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{

			"status_code": 419,

			"status_msg": "用户不存在",

			"user_id": 0,

			"token": "",
		})

		return

		//}

	}

	//判断密码是否正确
  // 密码校验   ---》
  
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {

		ctx.JSON(http.StatusUnprocessableEntity, gin.H{

			"status_code": 422,

			"status_msg": "password error",

			"user_id": 0,

			"token": "",
		})
		return

	}

	//返回结果
  token,err := util.GenerateToken(user.ID, user.Username)
	if err !=nil{
    ctx.JSON(500, gin.H{

			"status_code": 500,

			"status_msg": "服务器内部错误",

			"user_id": 0,

			"token": "",
		})
		return
  }
  ctx.JSON(http.StatusOK, gin.H{

		"status_code": 0,

		"status_msg": "success",

		"user_id": user.ID,

		"token": token,
	})

}
