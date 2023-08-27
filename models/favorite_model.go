package models

import "gorm.io/gorm"

// 点赞表
type Favorite struct {
	UserID  int64 `json:"user_id" gorm:"column:user_id; primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
	User    User  `json:"-" gorm:"foreignKey:UserID" `
	VideoID int64 `json:"video_id,omitempty" gorm:"column:video_id; primary_key" sql:"type:INT(10) UNSIGNED NOT NULL"`
	Video   Video `json:"-" gorm:"foreignKey:VideoID"`
}

func AddFavorite(userID int64, videoID int64) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&Favorite{}).Create(&Favorite{UserID: userID, VideoID: videoID}).Error; err != nil {
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ?", videoID).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func DelFavorite(userID int64, videoID int64) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("user_id = ? and video_id = ?", userID, videoID).Delete(&Favorite{}).Error; err != nil {
			return err
		}
		if err := tx.Model(&Video{}).Where("id = ?", videoID).
			UpdateColumn("favorite_count", gorm.Expr("favorite_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}
