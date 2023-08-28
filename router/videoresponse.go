package router

import (
	"douyin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResCode int64

const CodeSuccess ResCode = 0

var codeMsgMap = map[ResCode]string{
	CodeSuccess: "success",
}

type ResponseData struct {
	Code      ResCode     `json:"status_code"`
	Msg       interface{} `json:"status_msg"`
	VideoInfo interface{} `json:"video_list"`
}

func (rc ResCode) Msg() string {
	msg := codeMsgMap[rc]      //若错误返回空

	return msg
}

func ResponseSuccess(c *gin.Context, videoList []*model.VideoInfo) {
	c.JSON(http.StatusOK, &ResponseData{
		Code:      CodeSuccess,
		Msg:       CodeSuccess.Msg(),
		VideoInfo: &videoList, //返回的应该为数据库中的信息，该信息videoList用mysql/feed中的SelectVideoInfo()获得
	})
}
