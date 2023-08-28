package models

type Message struct {
	Id         int64  `json:"id"`
	Content    string `json:"content"`
	CreateTime string `json:"create_time"`
	FromUserId int64  `json:"from_user_id"`
	ToUserId   int64  `json:"to_user_id"`
}

func AddMessage(message *Message) error {
	return DB.Create(&message).Error
}

func GetMessages(userId int64, toUserId int64, messageList *[]*Message) error {
	return DB.Model(&Message{}).Where("from_user_id = ? and to_user_id = ?", userId, toUserId).Find(&messageList).Error
}
