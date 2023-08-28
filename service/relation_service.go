package service

import (
	"errors"
	"github.com/RaymondCode/simple-demo/models"
	"github.com/RaymondCode/simple-demo/utils"
	"strconv"
)

type Response models.Response
type FollowResponse models.FollowResponse

func (q *Response) RelationAction(userId int64, toUserId int64, actionType string) error {
	err := utils.ValidateActionType(actionType)
	if err != nil {
		return err
	}

	if have := models.IsUserExistById(strconv.FormatInt(toUserId, 10)); !have {
		return errors.New("用户不存在")
	}

	if toUserId == userId {
		return errors.New("无法关注自己")
	}

	//关注
	if actionType == "1" {
		follow := models.Follow{
			FollowUserId:   userId,
			FollowerUserId: toUserId,
		}

		if have := models.GetIsFollow(follow); have {
			return errors.New("已经关注过不能继续关注")
		}
		
		err = models.AddFollow(follow)
		if err != nil {
			return err
		}
	} else {
		follow := models.Follow{
			FollowUserId:   userId,
			FollowerUserId: toUserId,
		}
		err = models.DelFollow(follow)
		if err != nil {
			return err
		}
	}

	return nil
}

func (q *FollowResponse) GetFollowList(userId int64, key string, userKey string) error {
	q.UserList = []*models.User{}

	err := models.GetFollow(userId, &q.UserList, key, userKey)
	if err != nil {
		return err
	}
	return nil
}
func (q *FollowResponse) GetFriendList(userId int64) error {
	q.UserList = []*models.User{}
	err := models.GetFriend(userId, &q.UserList)
	if err != nil {
		return err
	}
	return nil
}
