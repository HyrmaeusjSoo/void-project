package service

import (
	"chat/internal/model"
	"chat/internal/repository/mysql"
	"chat/pkg/md5"
	"fmt"
	"math/rand"
	"time"
)

type UserService struct{}

var udb = &mysql.UserRepository{}

// 获取账号
func (*UserService) Fetch(id uint) (*model.User, error) {
	user, err := udb.GetById(id)
	user.SecureClear() //清除敏感信息
	return user, err
}

func (*UserService) List() ([]*model.User, error) {
	return udb.GetList()
}

// 账号是否存在
func (*UserService) ExistsAccount(account string) bool {
	return udb.ExistsAccount(account)
}

// 注册用户
func (*UserService) Register(user *model.User) error {
	salt := fmt.Sprintf("%d", rand.Int31())
	t := time.Now()
	user.Password = md5.SaltPassword(user.Password, salt)
	user.Salt = salt
	user.LoginTime = &t
	user.LoginOutTime = &t
	user.HeartBeatTime = &t
	err := udb.Create(user)
	user.SecureClear() //清除敏感信息
	return err
}

// 按账号获取账户
func (*UserService) GetByAccount(account string) (*model.User, error) {
	return udb.GetByAccount(account)
}

// 账号密码获取账户
func (*UserService) GetByAccountPassword(account, password string) (*model.User, error) {
	return udb.GetByAccountPassword(account, password)
}

// 修改用户信息
func (*UserService) Update(user *model.User) error {
	err := udb.Update(user)
	user.SecureClear()
	return err
}

// 删除用户
func (*UserService) Delete(id uint) error {
	return udb.Delete(id)
}
