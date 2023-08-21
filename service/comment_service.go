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

// 评论列表
func CommentList(videoID int64) (commentList []models.Comment, err error) {
	comments, err := models.QueryCommentByVideoID(videoID)
	if err != nil {
		return nil, err
	}
	commentIds := make([]int64, len(comments))
	for i, comments := range comments {
		commentIds[i] = comments.Id
		comment := models.Comment{}
		result := models.DB.Model(&models.Comment{}).Where("Id = ?", commentIds[i]).Find(&comment)
		if result == nil {
			return nil, result.Error
		}
		user := models.User{}
		res := models.DB.Model(&models.User{}).Where("Id = ?", comments.UserID).Find(&user)
		if res == nil {
			return nil, res.Error
		}
		comment.User = user
		commentList = append(commentList, comment)
	}
	return commentList, nil
}
