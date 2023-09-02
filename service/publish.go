package service

import (

	//"douyin-project/config"

	"net/http"
  "fmt"
	"github.com/22722679/douyin-project/model"
  "crypto/md5"
	"github.com/22722679/douyin-project/mysql"
  "gorm.io/gorm"
	"strconv"
)

func Register(request *model.ReAndLogRequest) model.ReAndLogResponse {
  var errID int64 = -1 
  perrId := &errID
  //check长度
  if len(request.Username) >32 || len(request.Password) >32 {
    msg := "account or password is too long"
    return model.ReAndLogResponse{
      StatusCode:    http.StatusBadRequest,
      StatusMsg:     &msg,
      UserID:        perrId,
      Token:         nil,
    }
  }
  fmt.Printf("%v\n", request);
  
  // 重复性检验
  count := mysql.CheckUserName(request.Username)
  if count != 0 {
    msg := "用户已存在"
    return model.ReAndLogResponse{
      StatusCode:    http.StatusBadRequest,
      StatusMsg:     &msg,
      UserID:        perrId,
      Token:         nil,
    }
  }
  //密码加密
  request.Password = fmt.Sprintf("%x", md5.Sum([]byte(request.Password)))
  //插入数据库
  regisInfo := model.RegisterInfo {
    Model:       gorm.Model{},
    UserName:    request.Username,
    Password:    request.Password,
  }
  //插入到注册信息表
  userId, _ := mysql.InserRegisterInfo(regisInfo) // 返回int64
  userInfo := model.User{
    Model: gorm.Model{ID: uint(userId),},
    Username: request.Username,
    Password: request.Password,
    
  }
  fmt.Printf("sss\n");
  //将用户插入用户信息表
  mysql.InsertUserInfo(userInfo)

  msg := "注册成功"
  return model.ReAndLogResponse{
      StatusCode:    http.StatusOK,
      StatusMsg:     &msg,
      UserID:        &userId,
      Token:         nil,
    }
}


func UserInfoService(userId uint) (model.UserInfoResponse, error) {

	userInfo, err := GetUserBaseInfo(userId, strconv.Itoa(int(userId)))

	if err != nil {

		return model.UserInfoResponse{}, err

	}



	//返回用户信息

	msg := "查询成功"

	return model.UserInfoResponse{

		StatusCode: 0,

		StatusMsg:  &msg,

		UserInfo:   userInfo,

	}, nil

}



func GetUserBaseInfo(tUserId uint, userId string) (model.UserInfo, error) {

	//查询用户名称是否存在

	userName, err := mysql.SelectUserName(int64(tUserId))

	if err != nil {

		return model.UserInfo{}, err

	}

	author := model.UserInfo{

		Id:            tUserId,

		FollowCount:   24321,

		FollowerCount: 23040,

		IsFollow:      true,

		Name:          userName,

		//Avatar:        config.Avatar,

		//BackgroundImages: config.BackgroundImages,

		//Signature:        config.Signature,

		//TotalFavorited:   config.TotalFavorited,

		//WorkCount:        config.WorkCount,

		//FavoriteCount:    config.FavoriteCountl,

	}

	return author, nil

}



