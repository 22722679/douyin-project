package router



import (

	"github.com/22722679/douyin-project/mysql"



	"github.com/22722679/douyin-project/controller"



	"github.com/gin-gonic/gin"

	//"net/http"

)



func InitRouter(r *gin.Engine) {
  //初始化ftp服务器连接
	//ftp.InitFTP()
	//--------------------------------

	r.GET("/douyin/feed", VideoInfoHandler)

	r.POST("/douyin/user/register/", controller.Register)

	r.POST("/douyin/user/login/", controller.Login)

	r.GET("/douyin/user", controller.UserInfo)

	r.POST("/douyin/publish/action/", controller.Publish)

	r.GET("/douyin/publish/list/", controller.PublishList)

	r.GET("/douyin/favorite/list/", controller.FavoriteList)
	
	r.POST("/douyin/favorite/action/",controller.FavoriteAction)
}



// 视频信息处理

func VideoInfoHandler(c *gin.Context) {

	data, _ := mysql.SelectVideoInfo()

	ResponseSuccess(c, data)

}

