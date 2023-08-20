package models

import (
	"gorm.io/gorm"
)

type FollowResponse struct {
	Response
	UserList []*User `json:"user_list"`
}

func AddFollow(follow Follow) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&follow).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", follow.FollowerUserId).UpdateColumn("follower_count", gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", follow.FollowUserId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

func DelFollow(follow Follow) error {
	err := DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("follow_userId = ? and follower_userId = ?", follow.FollowUserId, follow.FollowerUserId).Delete(&follow).Error; err != nil {
			return err
		}

		if err := tx.Model(&User{}).Where("id = ?", follow.FollowerUserId).UpdateColumn("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}
		if err := tx.Model(&User{}).Where("id = ?", follow.FollowUserId).UpdateColumn("follow_count", gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

func GetFollow(userId int64, followList *[]*User, key string, userKey string) error {
	res := DB.Model(&Follow{}).Select("users.*").Where(key+" = ?", userId).
		Joins("left join users on users.id = follows." + userKey).Find(&followList)
	return res.Error
}
func GetFriend(userId int64, followList *[]*User) error {
	var follows []*Follow
	res := DB.Model(&Follow{}).Select("follower_userId").Where("follow_userId = ?", userId).Find(&follows)
	if res.Error != nil {
		return res.Error
	}
	var followerIds []int64
	var newFollow *Follow
	for _, follow := range follows {
		//有结果的话证明是好友
		res1 := DB.Model(&Follow{}).Where("follow_userId = ? and follower_userId = ?", follow.FollowerUserId, userId).Find(&newFollow)
		if res1.Error != nil {
			return res1.Error
		}
		if res1.RowsAffected != 0 {
			followerIds = append(followerIds, follow.FollowerUserId)
		}
	}
	return DB.Model(&User{}).Where("id in ?", followerIds).Find(&followList).Error
}
func IsFriend(userId int64, toUserId int64) (bool, error) {
	var follow Follow
	res1 := DB.Model(&Follow{}).Where("follow_userId = ? and follower_userId = ?", userId, toUserId).First(&follow)
	res2 := DB.Model(&Follow{}).Where("follow_userId = ? and follower_userId = ?", toUserId, userId).First(&follow)
	if res1.Error != nil {
		return false, res1.Error
	}
	if res2.Error != nil {
		return false, res1.Error
	}
	if res2.RowsAffected != 0 && res1.RowsAffected != 0 {
		return true, nil
	}
	return false, nil
}
