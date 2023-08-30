package controller



import (

	"github.com/22722679/douyin-project/model"

	"github.com/22722679/douyin-project/service"

	"fmt"

	"net/http"

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



func PublishAction(ctx *gin.Context) {

	//参数解析

	request := new(model.PublishActionRequest)

	if err := ctx.ShouldBind(request); err != nil {

		fmt.Println("解析失败")

		return

	}



	userId, _ := ctx.Get("userId")

	code_a, _ := service.PublishAction(request, userId.(uint))

	ctx.JSON(http.StatusOK, &model.Response{

		StatusCode: code_a,

		StatusMsg:  model.GetMsg(code_a),

	})

	return

}

