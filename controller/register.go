package controller

import (
	//"database/sql"
	"douyin-project/model"
	"douyin-project/mysql"
	"fmt"
	"net/http"

	//"sync/atomic"
	"github.com/gin-gonic/gin"
	//"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

//var userIdSequence = int64(1)

func Register(ctx *gin.Context) {
	db := mysql.GETDB()

	//获取参数
	var requestUser model.User
	ctx.Bind(&requestUser)

	username := requestUser.Username
	password := requestUser.Password

	if len(username) != 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": 423,
			"status_msg":  "error",
			"user_id":     423,
			"token":       "用户不能为空",
		})
		return
	}
	if len(password) >= 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": 424,
			"status_msg":  "error",
			"user_id":     424,
			"token":       "密码不能少于6位",
		})
		return
	}

	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"status_code": 500,
			"status_msg":  "error",
			"user_id":     500,
			"token":       "密码加密错误",
		})
		return
	}
	newUser := model.User{
		Username: username,
		Password: string(hasedPassword),
	}

	sqlStr := fmt.Sprintf("insert into users(name,password) values(%s,%s)", newUser.Username, newUser.Password)
	db.Select(&newUser, sqlStr)

	//返回结果
	ctx.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "success",
		"user_id":     0,
		"token":       "注册成功",
	})

}
