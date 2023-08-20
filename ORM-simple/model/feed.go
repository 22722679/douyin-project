package model

//import "gorm.io/gorm"

type VideoInfo struct {
	//gorm.Model
	Id            int    `json:"id" db:"author_id"`
	PlayUrl       string `json:"play_url" db:"play_url"`
	CoverUrl      string `json:"cover_url" db:"cover_url"`
	FavoriteCount int    `json:"favorite_count" db:"favorite_count"`
	CommentCount  int    `json:"comment_count" db:"comment_count"`
	Title         string `json:"title" db:"title"`
}
