package controller

import (
	//"database/sql"
	"douyin-project/model"
	"douyin-project/mysql"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	db := mysql.GETDB()

	var requestUser model.User
	ctx.Bind(&requestUser)
	Name := requestUser.Username
	password := requestUser.Password

	//数据检验
	if len(Name) < 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": 421,
			"status_mag":  "账户必须为11位",
			"user_id":     421,
			"token":       "账户必须为11位",
		})
		return
	}
	if len(password) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": 420,
			"status_msg":  "密码不能少于6位",
			"user_id":     420,
			"token":       "密码不能少于6位",
		})
		return
	}

	//判断账户是否存在
	var user model.User

	sqlStr := "select name from users"
	if err := db.Select(&user, sqlStr); err != nil {
		//if err == sql.ErrNoRows {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": 419,
			"status_msg":  "用户不存在",
			"user_id":     419,
			"token":       "用户不存在",
		})
		return
		//}
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err == nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": 418,
			"status_msg":  "密码错误",
			"user_id":     418,
			"token":       "密码错误",
		})
	}

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_id":     0,
		"token":       "登录成功",
	})
}
