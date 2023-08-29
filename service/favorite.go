package service



import (

	"github.com/22722679/douyin-project/config"

	"github.com/22722679/douyin-project/model"

	"github.com/22722679/douyin-project/mysql"

	"strconv"

)



func FavoriteList(userId uint) (model.PublishListResponse, error) {

	//查询出视频id信息

	videoBaseList, err := mysql.SelectVideoInfoListByUserId(int64(userId))

	if err != nil {

		msg := "查询点赞视频id操作失败"

		return model.PublishListResponse{

			StatusCode: 1,

			StatusMsg:  &msg,

			UserList:   nil,

		}, err

	}



	videos := make([]model.Video, 0)



	for _, video := range videoBaseList {

		author, _ := GetUserBaseInfo(userId, strconv.Itoa(int(userId)))

		videos = append(videos, model.Video{

			Author:        author,

			CommentCount:  config.CommentCount,

			CoverUrl:      config.CoverUrl,

			FavoriteCount: config.FavoriteCount,

			Id:            int(video.Id),

			PlayUrl:       config.PlayUrl,

			Title:         config.Tiele,

		})

	}

	msg := "查询成功"

	return model.PublishListResponse{

		StatusCode: 0,

		StatusMsg:  &msg,

		UserList:   videos,

	}, nil

	// var favoriteList []model.Favorite

	// VideoList := make([]model.VideoInfo, 0)



	// if err := mysql.DB.Table("favorite").Where("user_id=? AND status=?", userId, 1).Find(&favoriteList).Error; err != nil {

	// 	return VideoList, nil

	// }

	// for _, e := range favoriteList {

	// 	var video = model.VideoInfo{}

	// 	if err := mysql.DB.Table("video_infos").Where("id=?", e.VideoId).Find(&video).Error; err != nil {

	// 		return nil, err

	// 	}

	// 	VideoList = append(VideoList, video)

	// }

	// return VideoList, nil

}
