package controller

import (
	"github.com/22722679/douyin-project/model"
	"github.com/22722679/douyin-project/service"

	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

// 获取user信息

func UserInfo(ctx *gin.Context) {

	//参数读取与解析

	userIs := ctx.Query("user_id")

	userId, _ := strconv.Atoi(userIs)



	response, err := service.UserInfoService(uint(userId))

	if err != nil {
		msg := "error"
		ctx.JSON(http.StatusOK, model.UserInfoResponse{
			StatusCode:  1,
			StatusMsg:   &msg,
		})
		return 
	}

	ctx.JSON(http.StatusOK, response)

	return

}

