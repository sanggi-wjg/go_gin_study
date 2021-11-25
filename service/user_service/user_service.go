package user_service

import (
	"errors"

	"github.nhnent.com/godo/cfo/models"
	"github.nhnent.com/godo/cfo/server/util"
)

var (
	errExistUsername = errors.New("이미 존재하는 username 입니다.")
	errWrongUser     = errors.New("wrong username or password 입니다.")
)

type User struct {
	UserName  string
	Password  string
	UserLevel uint
}

func (u *User) CheckUserByUserNameAndPassword() (bool, error) {
	res, err := models.CheckUserByUserNameAndPassword(u.UserName, u.Password)
	if err != nil {
		return false, err
	}
	if res != true {
		return false, errWrongUser
	}

	return true, nil
}

func (u *User) CreateUser() (uint, error) {
	isExistUser := models.IsExistUsername(u.UserName)
	if isExistUser {
		return 0, errExistUsername
	}

	u.Password, _ = util.HashPassword(u.Password)
	return models.CreateUser(u.UserName, u.Password, u.UserLevel)
}
