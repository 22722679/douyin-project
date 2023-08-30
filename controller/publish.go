package controller

import (

	//"database/sql"

	"github.com/22722679/douyin-project/model"
	"github.com/22722679/douyin-project/service"



	"net/http"


	"github.com/gin-gonic/gin"


)

func Register(ctx *gin.Context){
  request := new(model.ReAndLogRequest)
  request.Username,_ = ctx.GetQuery("username");
  request.Password,_ = ctx.GetQuery("password");
  // err := ctx.ShouldBindQuery(request)
  // fmt.Printf("%v\n", request); 
  // if err != nil {
  //   fmt.Println("解析失败")
  //   return
  // }
  response := service.Register(request)
  ctx.JSON(http.StatusOK,response)
  return
}
