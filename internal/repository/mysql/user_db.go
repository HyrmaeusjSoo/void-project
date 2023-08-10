package mysql

import (
	"errors"
	"void-project/internal/model"
	"void-project/internal/model/base"
	"void-project/internal/repository"
	"void-project/internal/repository/driver"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: driver.MySQL}
}

// 账户列表
func (u *UserRepository) GetList(pager base.Pager) ([]model.User, int, error) {
	var list []model.User
	total, err := repository.Paginate(
		u.db.Select("id", "account", "name", "avatar", "gender", "phone", "email", "is_login_out", "device_info"),
		&list, pager)
	if err != nil {
		return nil, 0, err
	}

	return list, int(total), nil
}

// 查询账户
func (u *UserRepository) GetById(id uint) (*model.User, error) {
	user := &model.User{}
	err := u.db.First(user, id).Error
	return user, err
}

// 账号是否存在
func (u *UserRepository) ExistsAccount(account string) bool {
	var count int64
	u.db.Model(model.User{}).Unscoped().Where("account = ?", account).Count(&count)
	return count >= 1
}

// 按账号查询账户
func (u *UserRepository) GetByAccount(account string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("account = ?", account).First(user).Error
	return user, err
}

// 账号密码查询账户
func (u *UserRepository) GetByAccountPassword(account, password string) (*model.User, error) {
	user := &model.User{}
	tx := u.db.Where("account = ? AND password = ?", account, password).First(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, errors.New("未找到用户")
	}
	return user, nil
}

// 账户列表in ids
func (u *UserRepository) GetInIds(ids []uint) ([]model.User, error) {
	var list []model.User
	err := u.db.Where("id IN ?", ids).Select("id", "account", "name", "avatar", "gender", "phone", "email", "is_login_out", "device_info").Find(&list).Error
	if err != nil {
		return list, err
	}
	return list, nil
}

// 新增用户
func (u *UserRepository) Create(user *model.User) error {
	tx := u.db.Create(user)
	if tx.RowsAffected == 0 {
		if tx.Error != nil {
			return tx.Error
		}
		return errors.New("新增用户失败")
	}
	return nil
}

// 更新账户
func (u *UserRepository) Update(user *model.User) error {
	tx := u.db.Model(user).Updates(model.User{
		Name:   user.Name,
		Avatar: user.Avatar,
		Gender: user.Gender,
		Phone:  user.Phone,
		Email:  user.Email,
	})
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("更新0条记录")
	}
	return nil
}

// 删除账户
func (u *UserRepository) Delete(id uint) error {
	if tx := u.db.Delete(&model.User{}, id); tx.RowsAffected == 0 {
		if tx.Error != nil {
			return tx.Error
		}
		return errors.New("删除0条记录")
	}
	return nil
}
