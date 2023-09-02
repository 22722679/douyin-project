package controller

import (
	"github.com/22722679/douyin-project/model"

	"github.com/22722679/douyin-project/service"

	"fmt"

	"net/http"

	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)



func PublishList(ctx *gin.Context) {

	//解析

	token := ctx.Query("token")

	userIdStr := ctx.Query("user_id")



	//获取当前状态登录用户的id、token

	loginuser, _ := strconv.Atoi(token)

	userId, _ := strconv.Atoi(userIdStr)



	if loginuser == 0 {

		response, err := service.PublishList(uint(userId), uint(userId))

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, response)

			return

		}

		ctx.JSON(http.StatusOK, response)

	} else {

		response, err := service.PublishList(uint(userId), uint(loginuser))

		if err != nil {

			ctx.JSON(http.StatusInternalServerError, response)

			return

		}

		ctx.JSON(http.StatusOK, response)

	}

	return

}

func Publish(ctx *gin.Context){
	//获取userId
	getUserId, _ := ctx.Get("user_id")
	var userId uint
	if re, err := getUserId.(uint); err {
		userId = re
	}
	//接收请求参数信息
	//title := ctx.PostForm("title")
	data, err := ctx.FormFile("data")
	if err != nil {
		ctx.JSON(http.StatusOK,&model.Response{
			StatusCode:  			1,
			StatusMsg:				err.Error(),
		})
		return
	}

	//返回至前端页面的展示信息
	filename := filepath.Base(data.Filename)
	finaname := fmt.Sprintf("%d_%s",userId,filename)

	//先存本地，再保存到云端，并获取云端地址
	save := filepath.Join("../videos/",finaname)
	if err := ctx.SaveUploadedFile(data,save); err != nil {
		ctx.JSON(http.StatusOK,model.Response{
			StatusCode:			1,
			StatusMsg:			err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK,model.Response{
		StatusCode:				0,
		StatusMsg:				finaname + "--upload success",
	})
}
