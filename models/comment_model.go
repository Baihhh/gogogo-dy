package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	// gorm.Model
	Id      int64  `json:"id" gorm:"column:id"`
	UserID  int64  `json:"-" gorm:"column:user_id"`
	User    User   `json:"user,omitempty" gorm:"foreignKey:UserID"`
	VideoID int64  `json:"-" gorm:"column:video_id"`
	Video   Video  `json:"-" gorm:"foreignKey:VideoID"`
	Content string `json:"content,omitempty" gorm:"column:content"`
	// CreateDate string `json:"create_date,omitempty"`
	CreatedAt time.Time `json:"created_date,omitempty"`
}

func AddComment(comment *Comment) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Comment{}).Create(comment).Error; err != nil {
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ?", comment.VideoID).
			UpdateColumn("comment_count", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func DelComment(comment *Comment) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Comment{}).Delete(comment).Error; err != nil {
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ?", comment.VideoID).
			UpdateColumn("comment_count", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func QueryCommentByID(comment *Comment, commentID int64) error {
	return DB.First(comment, commentID).Error
}
