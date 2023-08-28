package models

type User struct {
	Id              int64  `json:"id" gorm:"column:id"`
	Name            string `json:"name" gorm:"column:name"`
	Password        string `json:"-" gorm:"column:password"`
	Avatar          string `json:"avatar" gorm:"column:avatar"`
	BackgroundImage string `json:"background_image" gorm:"column:background_image"`
	Signature       string `json:"signature" gorm:"column:signature"`
	FollowCount     int64  `json:"follow_count" gorm:"default:0"`
	FollowerCount   int64  `json:"follower_count" gorm:"default:0"`
	IsFollow        bool   `json:"is_follow" gorm:"-"`
	TotalFavorited  string `json:"total_favorited" gorm:"default:0"`
	WorkCount       int64  `json:"work_count" gorm:"default:0"`
	FavoriteCount   int64  `json:"favorite_count" gorm:"default:0"`
}

func AddUser(user *User) error {
	return DB.Create(&user).Error
}
func IsUserExistByUsername(username string) bool {
	user, ok := QueryUserLogin(username, "name")
	if !ok || user.Id == 0 {
		return false
	}
	return true
}

func IsUserExistById(id string) bool {
	user, ok := QueryUserLogin(id, "id")
	if !ok || user.Id == 0 {
		return false
	}
	return true
}

func QueryUserLogin(username string, key string) (User, bool) {
	var user User
	res := DB.Table("users").Where(key+" = ?", username).First(&user)
	if res.Error != nil || res.RowsAffected == 0 {
		return User{}, false
	}
	return user, true
}

func QueryIsFollow(userId int64, myUserId int64) bool {
	var follow Follow
	res := DB.Table("follows").Where("follow_userId = ? and follower_userId = ?", myUserId, userId).First(&follow)
	if res.Error != nil || res.RowsAffected == 0 {
		return false
	}
	return true
}
