package models

import (
	"gorm.io/gorm"
	"time"
	// "gorm.io/gorm"
)

type Video struct {
	// gorm.Model // 包含Id、CreatedAt、UpdatedAt
	Id            int64     `json:"id" gorm:"column:id"`
	Title         string    `json:"title" gorm:"column:title"`
	AuthorID      int64     `json:"-" gorm:"column:author_id"`
	Author        User      `json:"author" gorm:"foreignKey:AuthorID"`
	PlayUrl       string    `json:"play_url" gorm:"column:paly_url"`
	CoverUrl      string    `json:"cover_url" gorm:"column:cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite" gorm:"-"`
	CreatedAt     time.Time `json:"-" gorm:"column:created_at"`
	// UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
}

func AddVideo(video *Video) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&video).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", video.AuthorID).UpdateColumn("work_count", gorm.Expr("work_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}
