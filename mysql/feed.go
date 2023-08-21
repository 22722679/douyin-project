package mysql

import (
	"database/sql"
	"douyin-project/model"
	"go.uber.org/zap"
)

func SelectVideoInfo() (videoList []*model.VideoInfo, err error) {
	sqlStr := "select author_id, play_url, cover_url, favorite_count, comment_count, title from video_infos"
	if err := db.Select(&videoList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("数据库中没有视频")
			err = nil
		}
	}
	return
}
