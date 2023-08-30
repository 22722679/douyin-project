package controller

import (
//	"log"

	"github.com/22722679/douyin-project/service"
 // "github.com/22722679/douyin-project/model"

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
func favoriteAction(ctx *gin.Context){
  // //判断用户是否登录
  // stUserId := ctx.GetString("userId")
  // userId, _ := strconv.ParseInt(stUserId,10,64)
  // strVideoId := ctx.Query("video_id")
  // videoId, _ := strconv.ParseInt(strVideoId,10,64)
  // strActionType := ctx.Query("action_type")
  // actionType, _ := strconv.ParseInt(strActionType,10,64)
  // like := new(service.LikeService)

  // //点赞操作
  // err := like.FavouriteAction(userId,videoId,int32(actionType))
  // if err == nil {
  //   log.Printf("favoriteAction获取userid，videoId，actiontype成功")
  //   ctx.JSON(http.StatusOK,likeResponse{
  //       StatusCode:       0,
  //       StatusMsg:       "favorite action success",
  //   })
  // } else {
  //   log.Printf("Action获取userid，videoId，actiontype失败：%v",err)
  //   ctx.JSON(http.StatusOK,likeResponse{
  //     StatusCode:         1,
  //     StatusMsg:          "favorite action error",
  //   })
  // }
}

type likeResponse struct {
  StatusCode    int32     `json:"status_code"`
  StatusMsg     string    `json:"status_msg,omitempty"`
}

