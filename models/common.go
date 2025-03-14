package models

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User User `json:"user,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []*Video `json:"video_list"`
	NextTime  int64    `json:"next_time"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"content,omitempty"`
}

type MessageResponse struct {
	Response
	MessageList []*Message `json:"message_list"`
}

type Follow struct {
	//关注者
	FollowUserId int64 `json:"follow_userId,omitempty" gorm:"column:follow_userId"`
	//被关注者
	FollowerUserId int64 `json:"follower_userId,omitempty" gorm:"column:follower_userId"`
}
