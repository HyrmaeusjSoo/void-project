package mysql

import (
	"chat/internal/model"
	"chat/internal/repository/driver"
	"errors"
)

type UserRepository struct{}

// 账户列表
func (*UserRepository) GetList() ([]*model.User, error) {
	var list []*model.User
	driver.MySQL.Select("account", "name", "avatar", "gender", "phone", "email", "is_login_out", "device_info").Find(&list)
	/* if tx := driver.MySQL.Find(&list); tx.RowsAffected == 0 {
		return nil, errors.New("获取用户列表失败")
	} */
	return list, nil
}

// 查询账户
func (*UserRepository) GetById(id uint) (*model.User, error) {
	user := &model.User{}
	if tx := driver.MySQL.First(user, id); tx.RowsAffected <= 0 {
		return nil, errors.New("未找到用户！")
	}
	return user, nil
}

// 账号是否存在
func (*UserRepository) ExistsAccount(account string) bool {
	var count int64
	driver.MySQL.Model(model.User{}).Where("account = ?", account).Count(&count)
	return count >= 1
}

// 按账号查询账户
func (*UserRepository) GetByAccount(account string) (*model.User, error) {
	user := &model.User{}
	tx := driver.MySQL.Where("account = ?", account).First(user)
	if tx.RowsAffected == 0 {
		return nil, errors.New("未找到用户")
	}
	return user, nil
}

// 账号密码查询账户
func (*UserRepository) GetByAccountPassword(account, password string) (*model.User, error) {
	user := &model.User{}
	tx := driver.MySQL.Where("account = ? AND password = ?", account, password).First(user)
	if tx.RowsAffected == 0 {
		return nil, errors.New("未找到用户")
	}
	return user, nil
}

// 新增用户
func (*UserRepository) Create(user *model.User) error {
	tx := driver.MySQL.Create(user)
	if tx.RowsAffected == 0 {
		return errors.New("新增用户失败")
	}
	return nil
}

// 更新账户
func (*UserRepository) Update(user *model.User) error {
	tx := driver.MySQL.Model(user).Updates(model.User{
		Name:   user.Name,
		Avatar: user.Avatar,
		Gender: user.Gender,
		Phone:  user.Phone,
		Email:  user.Email,
	})
	if tx.RowsAffected == 0 {
		return errors.New("更新用户失败")
	}
	return nil
}

// 删除账户
func (*UserRepository) Delete(id uint) error {
	if tx := driver.MySQL.Delete(&model.User{}, id); tx.RowsAffected == 0 {
		return errors.New("删除用户失败")
	}
	return nil
}
