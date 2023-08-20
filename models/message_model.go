package models

type Message struct {
	Id         int64  `json:"id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
}

func AddMessage(message *Message) error {
	return DB.Create(&message).Error
}

func GetMessages(userId int64, toUserId int64, messageList *[]*Message) error {
	return DB.Model(&Message{}).Where("user_id = ? and to_user_id = ?", userId, toUserId).Find(&messageList).Error
}
