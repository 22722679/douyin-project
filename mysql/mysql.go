包mysql

导入（
	“ORM-简单/配置”douyin-project
	“FMMT”

	_ “github.com/go-sql-driver/mysql”
	“github.com/jmoiron/sqlx”
	“github.com/spf13/viper”
	“go.uber.org/zap”
）

var db *sqlx 。D B

func Init() (err错误) {
	dsn := fmt 。Sprintf( "%s:%s@tcp(%s)/douyin?charset=utf8mb4&parseTime=True&loc=Local" , 配置.用户, 配置.密码, 配置.IpMessage )
	db, 错误 = sqlx 。连接（“mysql”，dsn）
	如果错误！= nil {
		扎普。L() 。Error( "连接数据库失败,错误码为：%v\n" , zap .Error (err))
		返回
	}
	数据库。SetMaxOpenConns(viper . GetInt( "mysql.max_open_conns" )) //设置数据库的最大打开连接数
	数据库。SetMaxIdleConns(viper . GetInt( "mysql.max_idle_conns" )) //设置连接空闲的最大时间，过期的连接可能会在关闭之前重用。
	返回
}

func Close() { //数据库关闭返回空
	_ = 数据库。关闭（）
}
