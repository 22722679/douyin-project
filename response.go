package包主

import进口 （
	“ORM-简单/模型”"douyin-project/model"
	“网络/http”"net/http"

	“github.com/gin-gonic/gin”"github.com/gin-gonic/gin"
）

type类型 ResCode int64int64

constconst CodeSuccess ResCode = 00

varvar codeMsgMap = 地图[ResCode]字符串{map[ResCode]string{
	代码成功：“成功”，"success",
}

type类型响应数据结构{struct {
	代码 ResCode `json:"status_code"``json:"status_code"`
	消息接口{} `json:"status_msg"`interface{} `json:"status_msg"`
	VideoInfo 接口{} `json:"video_list"`interface{} `json:"video_list"`
}

funcfunc (rc ResCode) Msg() 字符串 {string {
	msg, _ := codeMsgMap[rc]

	返回消息return msg
}

funcfunc ResponseSuccess(c *gin.Context, videoList []*model.VideoInfo) {.Context, videoList []*model.VideoInfo) {
	c.JSON(http.StatusOK, &ResponseData{.JSON(http.StatusOK, &ResponseData{
		代码：代码成功，
		消息：CodeSuccess.Msg(),.Msg(),
		VideoInfo: &videoList, //返回的应该为数据库中的信息，该信息videoList用mysql/feed中的SelectVideoInfo()获得//返回的应该为数据库中的信息，该信息videoList用mysql/feed中的SelectVideoInfo()获得
	})
}
