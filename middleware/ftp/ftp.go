package ftp

import (
	"douyin-project/config"
	"log"
	"time"

	"github.com/dutchcoders/goftp"
)

var MyFTP *goftp.FTP

func InitFTP() {
	//获取ftp链接
	var err error
	MyFTP, err = goftp.Connect(config.FtpIp)
	if err != nil {
		log.Printf("获取到了FTP链接失败!!!")
	}
	log.Printf("获取到FTP链接成功%v:", MyFTP)
	//登录
	err = MyFTP.Login(config.FtpName, config.FtpPass)
	if err != nil {
		log.Printf("FTP登录失败!!!")
	}
	log.Printf("FTP登录成功!!!")
	//维持长链接
	go keepAlive()
}

func keepAlive() {
	//如何维持长链接
	time.Sleep(time.Duration(config.HeartTime) * time.Second)
	MyFTP.Noop()
}
