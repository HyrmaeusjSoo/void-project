package service

import (
	"chat/internal/model"
	"chat/internal/repository/mysql"
)

type UserService struct{}

var udb = &mysql.User{}

// 获取账号
func (*UserService) Fetch(id uint) (any, error) {
	return udb.GetById(id)
}

// 账号是否存在
func (*UserService) ExistsAccount(account string) bool {
	return udb.ExistsAccount(account)
}

// 注册用户
func (*UserService) Register(user *model.User) error {
	return udb.CreateUser(user)
}

// 按账号获取账户
func (*UserService) GetByAccount(account string) (*model.User, error) {
	return udb.GetByAccount(account)
}

// 账号密码获取账户
func (*UserService) GetByAccountPassword(account, password string) (*model.User, error) {
	return udb.GetByAccountPassword(account, password)
}
