package main



import (

	"github.com/22722679/douyin-project/config"

	"github.com/22722679/douyin-project/mysql"

	"github.com/22722679/douyin-project/router"

	"fmt"



	"github.com/gin-gonic/gin"

)



func main() {

	r := gin.Default() //创建服务



	//初始化ftp服务器连接

	//ftp.InitFTP()

	//--------------------------------



	if err := mysql.Init(); err != nil { //数据库启动失败的返回

		fmt.Printf("mysql启动失败，err:%v\n", err)

		return

	}

	defer mysql.Close() //defer表示阻塞,运行到最后时，才会执行命令，要关闭数据库，所以最后关闭

	router.InitRouter(r)

	r.Run(fmt.Sprintf(":%d", config.Port))

}

