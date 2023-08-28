package model


//响应方法结构体
type Response struct {
	StatusCode     int32    `json:"status_code"`
	StatusMsg	   string   `json:"status_msg"`
}


//业务状态码
var errCode = map[int32]string{
	0:					"OK",				//成功
	500:				"error",			//失败
	1000:				"请您先登录账户",    //登录返回请求
	400:				"无权限",			//permission
	300:				"无效的token",      //InvaildToken
}

func GetMsg(c int32) string {
	return errCode[c]
}