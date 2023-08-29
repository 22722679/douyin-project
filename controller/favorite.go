package controller



import (

	"github.com/22722679/douyin-project/service"

	"net/http"

	"strconv"



	"github.com/gin-gonic/gin"

)



//视频列表查询

func FavoriteList(ctx *gin.Context){

	strUserId := ctx.Query("user_id")

	userId, _ := strconv.Atoi(strUserId)

	res, err := service.FavoriteList(uint(userId))



	if err != nil {

		ctx.JSON(http.StatusOK,res)

		return 

	}

	ctx.JSON(http.StatusOK,res)

	return

}
