package service

import (
	"github.com/RaymondCode/simple-demo/models"
)

func AddComment(userID int64, videoID int64, text string) *models.Comment {
	comment := &models.Comment{
		UserID:  userID,
		VideoID: videoID,
		Content: text,
	}
	models.AddComment(comment)

	return comment
}

func DelComment(commentID int64) error {
	comment := &models.Comment{}
	models.QueryCommentByID(comment, commentID)

	return models.DelComment(comment)
}
