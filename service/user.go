package service

import (
	//"douyin-project/config"
	"douyin-project/model"
	"douyin-project/mysql"
	"strconv"
)

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
