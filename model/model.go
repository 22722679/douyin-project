package model

//响应方法结构体

type Response struct {

	StatusCode     int32    `json:"status_code"`

	StatusMsg	   string   `json:"status_msg"`

}

//注册请求参数
type ReAndLogRequest struct{
  Username string `json:"username" db:"user_name"`  //用户名  
  Password string `json:"password" db:"password"`  //密码
}

type ReAndLogResponse struct {
  StatusCode    int64     `json:"status_code"`    //状态码
  StatusMsg    *string   `json:"status_msg"`     //状态描述
  Token        *string   `json:"token"`         // 用户鉴权token
  UserID       *int64    `json:"user_id"`        //用户id
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
