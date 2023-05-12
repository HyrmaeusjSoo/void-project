package service

import (
	"chat/internal/model"
	"chat/internal/repository/mysql"
)

type UserService struct{}

func (*UserService) Fetch(id uint) (any, error) {
	udb := mysql.User{}
	res, err := udb.GetById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (*UserService) ExistsAccount(account string) bool {
	return (&mysql.User{}).ExistsAccount(account)
}

func (*UserService) Register(user *model.User) error {
	return (&mysql.User{}).CreateUser(user)
}

func (*UserService) GetByAccount(account string) (*model.User, error) {
	return nil, nil
}
