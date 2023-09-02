package service



import (

	"github.com/22722679/douyin-project/model"

	"strconv"

	"github.com/22722679/douyin-project/config"

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



