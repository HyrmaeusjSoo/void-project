package service

import (
	"fmt"
	"math/rand"
	"time"
	"void-project/internal/model"
	"void-project/internal/repository/mysql"
	"void-project/pkg/md5"
)

type UserService struct {
	db *mysql.UserRepository
}

func NewUserService() *UserService {
	return &UserService{db: mysql.NewUserRepository()}
}

// 获取账号
func (u *UserService) Fetch(id uint) (*model.User, error) {
	user, err := u.db.GetById(id)
	user.SecureClear() //清除敏感信息
	return user, err
}

func (u *UserService) List(page, size int) ([]model.User, int, error) {
	return u.db.GetList(page, size)
}

// 账号是否存在
func (u *UserService) ExistsAccount(account string) bool {
	return u.db.ExistsAccount(account)
}

// 注册用户
func (u *UserService) Register(user *model.User) error {
	salt, t := fmt.Sprintf("%d", rand.Int31()), time.Now()
	user.Password = md5.SaltPassword(user.Password, salt)
	user.Salt = salt
	user.LoginTime = &t
	user.LoginOutTime = &t
	user.HeartBeatTime = &t
	err := u.db.Create(user)
	user.SecureClear() //清除敏感信息
	return err
}

// 按账号获取账户
func (u *UserService) GetByAccount(account string) (*model.User, error) {
	return u.db.GetByAccount(account)
}

// 账号密码获取账户
func (u *UserService) GetByAccountPassword(account, password string) (*model.User, error) {
	return u.db.GetByAccountPassword(account, password)
}

// 修改用户信息
func (u *UserService) Update(user *model.User) error {
	err := u.db.Update(user)
	user.SecureClear()
	return err
}

// 删除用户
func (u *UserService) Delete(id uint) error {
	return u.db.Delete(id)
}
