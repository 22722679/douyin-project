package controller

import (

	"github.com/22722679/douyin-project/model"
	"github.com/22722679/douyin-project/mysql"
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

		ctx.JSON(http.StatusInternalServerError,res)

		return 
	}

	ctx.JSON(http.StatusOK,res)

	return

}


//视频点赞功能
func FavoriteAction(ctx *gin.Context){
  //user_id参数与赞操作绑定
  getUserId, _ := ctx.Get("user_id")
  var userId uint
  if cs, err := getUserId.(uint); err{
    userId = cs
  }

  //参数获取
  StractionType := ctx.Query("action_type")
  actionType, _ := strconv.ParseUint(StractionType,10,10)
  StrVideoId := ctx.Query("video_id")
  VideoId, _ := strconv.ParseUint(StrVideoId,10,10)

  //service以及响应
  err := mysql.FavoriteAction(userId ,uint(VideoId),uint(actionType))
  if err != nil {
    ctx.JSON(http.StatusBadRequest,model.Response{
      StatusCode:       1,
      StatusMsg:        err.Error(),
    })
  } else{
    ctx.JSON(http.StatusOK,model.Response{
      StatusCode:      0,
      StatusMsg:       "success",
    })
  }
}


