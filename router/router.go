package router

import (
	"douyin-project/mysql"

	"douyin-project/controller"

	"github.com/gin-gonic/gin"
	//"net/http"
)

func InitRouter(r *gin.Engine) {
	r.GET("/douyin/feed", VideoInfoHandler)
	r.POST("/douyin/user/register/", controller.Register)
	r.POST("/douyin/user/login/", controller.Login)
	r.GET("/douyin/user", controller.UserInfo)
	r.POST("/douyin/publish/action/", controller.PublishAction)
	r.GET("/douyin/publish/list/", controller.PublishList)
	r.GET("/douyin/favorite/list/", controller.FavoriteList)
}

// 视频信息处理
func VideoInfoHandler(c *gin.Context) {
	data, _ := mysql.SelectVideoInfo()
	ResponseSuccess(c, data)
}
