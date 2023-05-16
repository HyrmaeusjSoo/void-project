package mysql

import (
	"chat/internal/model"
	"chat/internal/repository/db"
	"errors"
)

type User struct{}

// 查询账户
func (*User) GetById(id uint) (any, error) {
	user := &model.User{}
	if tx := db.MySQL.First(user, id); tx.RowsAffected <= 0 {
		return nil, errors.New("未找到用户！")
	}
	return user, nil
}

// 账号是否存在
func (*User) ExistsAccount(account string) bool {
	var count int64
	db.MySQL.Model(model.User{}).Where("account = ?", account).Count(&count)
	return count >= 1
}

// 新增用户
func (*User) CreateUser(user *model.User) error {
	tx := db.MySQL.Create(user)
	if tx.RowsAffected == 0 {
		return errors.New("新增用户失败")
	}
	return nil
}

// 按账号查询账户
func (*User) GetByAccount(account string) (*model.User, error) {
	user := &model.User{}
	tx := db.MySQL.Where("account = ?", account).First(user)
	if tx.RowsAffected == 0 {
		return nil, errors.New("未找到用户")
	}
	return user, nil
}

// 账号密码查询账户
func (*User) GetByAccountPassword(account, password string) (*model.User, error) {
	user := &model.User{}
	tx := db.MySQL.Where("account = ? AND password = ?", account, password).First(user)
	if tx.RowsAffected == 0 {
		return nil, errors.New("未找到用户")
	}
	return user, nil
}
