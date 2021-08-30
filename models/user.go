package models

import (
	"gorm.io/gorm"
)

type User struct {
	Idx          uint   `gorm:"primaryKey;autoIncrement;"`
	UserId       string `gorm:"size:20;unique;not null;"`
	UserEmail    string `gorm:"size:50;unique;not null;"`
	UserPassword string `gorm:"size:100;not null;"`
	Created      uint64 `gorm:"autoCreateTime"`
	LastLogin    uint64 `gorm:"default:0"`
	IsSuperUser  int    `gorm:"default:1"`
}

func IsAuthenticate(userId string, password string) (bool, error) {
	var user User
	err := db.Select("idx").Where(User{
		UserId: userId, UserPassword: password,
	}).First(&user).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	return true, nil
}

func GetUserByUserId(userId string) (*User, error) {
	var user User

	result := db.Where(&User{UserId: userId}).First(&user)
	err := result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &user, nil
}

func GetUserAll() ([]*User, error) {
	var users []*User

	result := db.Find(&users)
	err := result.Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

func CreateNormalUser(userId, userEmail, password string) (uint, error) {
	user := User{
		UserId:       userId,
		UserEmail:    userEmail,
		UserPassword: password,
	}
	result := db.Create(&user)
	err := result.Error
	if err != nil {
		return 0, err
	}

	return user.Idx, nil
}

func CreateSuperUser() {

}
