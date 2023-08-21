package main

import (
	"douyin-project/middleware/ftp"
	"douyin-project/mysql"
	"fmt"

	//"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	//初始化FTP服务器连接
	ftp.InitFTP()
	//--------------------
	if err := mysql.Init(); err != nil { //数据库启动失败的返回
		fmt.Printf("mysql启动失败，err:%v\n", err)
		return
	}
	defer mysql.Close() //defer表示阻塞,运行到最后时，才会执行命令，要关闭数据库，所以最后关闭
	r := gin.Default()  //创建服务
	r.GET("/douyin/feed", VideoInfoHandler)
	r.Run(":8080")
}

func VideoInfoHandler(c *gin.Context) {
	data, _ := mysql.SelectVideoInfo()
	ResponseSuccess(c, data)
}

// response层
//
