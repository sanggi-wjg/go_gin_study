package models

import (
	"errors"

	"github.nhnent.com/godo/cfo/server/util"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement;"`
	UserName  string `gorm:"size:50;not null;unique;"`
	Password  string `gorm:"size:200;not null;"`
	UserLevel uint   `gorm:"not null;default:3"`
	gorm.Model
}

var (
	errWrongUsername = errors.New("username 확인해주세요.")
	errWrongPasswd   = errors.New("password 확인해주세요.")
)

func CheckUserByUserNameAndPassword(username string, password string) (bool, error) {
	var user User
	err := db.Select("id", "password").Where(User{UserName: username}).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, errWrongUsername
		}
		return false, err
	}

	isValidPwd := util.CheckPasswordHash(password, user.Password)
	if !isValidPwd {
		return false, errWrongPasswd
	}
	// err = db.Select("id").Where(User{UserName: username, Password: pwd}).Take(&user).Error
	// if err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return false, errNotExistUser
	// 	}
	// 	return false, err
	// }
	if user.ID > 0 {
		return true, nil
	}
	return false, nil
}

func IsExistUsername(username string) bool {
	var user User
	res := db.Where("user_name = ?", username).First(&user)
	if res.Error == gorm.ErrRecordNotFound {
		return false
	}
	if user.ID > 0 {
		return true
	}
	return false
}

func CreateUser(username string, password string, userlevel uint) (uint, error) {
	user := User{UserName: username, Password: password, UserLevel: userlevel}
	res := db.Select("UserName", "Password", "UserLevel").Create(&user)
	if res.Error != nil {
		return 0, res.Error
	}

	return user.ID, nil
}
