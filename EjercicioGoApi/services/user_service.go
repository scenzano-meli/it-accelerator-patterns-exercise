package services

import (
	"../domains"
	"../utils"
)


func GetUser(userID int64) (*domains.User, *utils.ApiError) {
	user := &domains.User{
		Id: userID,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return user, nil
}