package service

import "go_gin_study/models"

type User struct {
	Idx          uint
	UserId       string
	UserPassword string
}

func (u *User) GetAll() ([]*models.User, error) {
	//var users []*models.User
	users, err := models.GetUserAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}
