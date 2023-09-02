package mysql



import (

	"github.com/22722679/douyin-project/config"

	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"

)



var db *sqlx.DB





func Init() (err error) {

	db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s)/douyin?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.PassWord, config.IpMessage))

	if err != nil {

		zap.L().Error("连接数据库失败,错误码为：%v\n", zap.Error(err))

		return

	}


	db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns")) //设置数据库的最大打开连接数

	db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns")) //设置连接空闲的最大时间,过期的连接可能会在重用之前惰性关闭。

	return

}



func Close() { //数据库关闭返回空

	_ = db.Close()

}


func GETDB() *sqlx.DB {

	return db

}
