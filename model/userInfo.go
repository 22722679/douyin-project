package model

//用户相关详细信息
type UserInfo struct {
	FollowCount      int64  `json:"follow_count" db:"follow_count"`     //关注总数
	FollowerCount    int64  `json:"follower_count" db:"follower_count"` //粉丝总数
	Id               uint   `json:"id" db:"id"`                         //用户id
	IsFollow         bool   `json:"is_follow" db:"is_follow"`           //是否被关注
	Name             string `json:"name" db:"name"`                     //用户名称
	Avatar           string `json:"avator"`                             //用户头像
	BackgroundImages string `json:"background_images"`                  //用户个人页面顶部大图
	Signature        string `json:"signature"`                          //个人简介
	TotalFavorited   string `json:"total_favorited"`                    //获赞数量
	//WorkCount        int64  `json:"work_count"`                         //作品数
	//FavoriteCount    int64  `json:"favoritecount"`                      //喜欢数
}

//user响应
type UserInfoResponse struct {
	StatusCode int64    `json:"status_code"` //状态码
	StatusMsg  *string  `json:"status_msg"`  //返回状态描述
	UserInfo   UserInfo `json:"user"`        //用户信息
}
