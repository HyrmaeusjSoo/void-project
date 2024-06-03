package mysql

import (
	"errors"
	"void-project/internal/model"
	"void-project/internal/model/base"
	"void-project/internal/repository"
	"void-project/internal/repository/driver"
	"void-project/internal/repository/mysql/scope"

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

// 统计账号
func (u *UserRepository) CountByAccount(account string) (count int64) {
	u.db.Model(model.User{}).Unscoped().Where("account = ?", account).Count(&count)
	return count
}

// 统计账号密码
func (u *UserRepository) CountByAccountPassword(account, password string) (count int64) {
	u.db.Model(model.User{}).Unscoped().Where("account = ? AND password = ?", account, password).Count(&count)
	return count
}

// 按账号查询账户
func (u *UserRepository) GetByAccount(account string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("account = ?", account).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未找到用户")
	}
	return user, err
}

// 账号密码查询账户
func (u *UserRepository) GetByAccountPassword(account, password string) (*model.User, error) {
	user := &model.User{}
	err := u.db.Where("account = ? AND password = ?", account, password).First(user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("未找到用户")
	}
	return user, err
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
	return scope.Create(u.db, user)
}

// 更新账户
func (u *UserRepository) Update(user *model.User) error {
	return scope.Update(u.db, user)
}

// 删除账户
func (u *UserRepository) Delete(id uint) error {
	return scope.Delete(u.db, &model.User{}, id)
}
