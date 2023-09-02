package mysql



import (

	"database/sql"

	"github.com/22722679/douyin-project/model"

	"fmt"

	"go.uber.org/zap"

)



func SelectVideoInfo() (videoList []*model.VideoInfo, err error) {

	sqlStr := "select author_id, play_url, cover_url, favorite_count, comment_count, title from video_infos"

	if err := db.Select(&videoList, sqlStr); err != nil {

		if err == sql.ErrNoRows {

			zap.L().Warn("数据库中没有视频")

		}

	}

	return

}



func SelectVideoInfoListByUserId(id int64) ([]*model.Video, error) {

	var UserId []*model.Video

	VideoList := fmt.Sprintf("select author_id from video_infos where id = %d", id)

	if err := db.Select(&UserId, VideoList);err != nil {

		if err == sql.ErrNoRows {

			zap.L().Warn("没有该用户")

		}

	}

	return UserId,nil

}
