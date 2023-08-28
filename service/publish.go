package service

import (
	"douyin-project/model"
	"strconv"
	"douyin-project/config"
)

func PublishList(userId, loginuserId uint) (model.PublishListResponse, error) {
	//用user_id查询视频信息
	// 用户基本信息
	author, err := GetUserBaseInfo(userId, strconv.Itoa(int(userId)))
	msg := "查询用基本信息失败"
	if err != nil {
		return model.PublishListResponse{
			StatusCode: 1,
			StatusMsg:  &msg,
			UserList:   nil,
		}, err
	}

	//视频基本信息
	//videoListBase, err := mysql.SelectVideoInfoListByUserId(int64(userId))
	//if err != nil {
	//	msg := "查询用户信息失败"
	//	return model.PublishListResponse{
	//		StatusCode: 1,
	//		StatusMsg: &msg,
	//		UserList: nil,
	//	},err
	//}

	//视频数据进行封装
	videos := make([]model.Video, 0)
	videos = append(videos, model.Video{
		Author:        author,
		CommentCount:  config.CommentCount,
		CoverUrl:      config.CoverUrl,
		FavoriteCount: config.FavoriteCount,
		Id:            config.Id,
		PlayUrl:       config.PlayUrl,
		Title:         config.Tiele,
	})

	msg_a := "查询成功"
	return model.PublishListResponse{
		StatusCode: 0,
		StatusMsg:  &msg_a,
		UserList:   videos,
	}, nil
}

func PublishAction(request *model.PublishActionRequest, userId uint) (int32, error) {
	//requestA := new(model.PublishActionRequest)
	//if err := ctx.
	return 0, nil
}
