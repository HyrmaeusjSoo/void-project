package service

import (
	"errors"
	"math/rand/v2"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"void-project/internal/model"
	"void-project/internal/model/base"
	"void-project/internal/repository/mysql"
	"void-project/pkg"
	"void-project/pkg/bcrypt"
	"void-project/pkg/logger"
	"void-project/pkg/md5"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserService struct {
	db *mysql.UserRepository
}

func NewUserService() *UserService {
	return &UserService{db: mysql.NewUserRepository()}
}

// 获取用户
func (u *UserService) Fetch(id uint) (*model.User, error) {
	user, err := u.db.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}
	if user != nil {
		user.SecureClear() //清除敏感信息
	}
	return user, err
}

// 获取用户列表
func (u *UserService) List(pager base.Pager) ([]model.User, int, error) {
	return u.db.GetList(pager)
}

// 账号是否存在
func (u *UserService) ExistsAccount(account string) bool {
	return u.db.CountByAccount(account) > 0
}

// 注册用户
func (u *UserService) Register(user *model.User) error {
	t := base.NewTime(time.Now())
	// salt := fmt.Sprintf("%d", rand.Int31())
	// user.Password = md5.SaltPassword(user.Password, salt)
	password, err := bcrypt.GeneratePassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	user.Identity = nil
	// user.Salt = &salt
	user.LoginTime = t
	user.LoginOutTime = t
	user.HeartBeatTime = t
	err = u.db.Create(user)
	if err != nil {
		return err
	}
	user.SecureClear() //清除敏感信息
	return nil
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
	err := u.db.Update(&model.User{
		Model: base.Model{
			ID: user.ID,
		},
		Name:   user.Name,
		Avatar: user.Avatar,
		Gender: user.Gender,
		Phone:  user.Phone,
		Email:  user.Email,
	})
	user.SecureClear()
	return err
}

// 修改账户密码
func (u *UserService) UpdatePassword(userId uint, password string) error {
	hashPwd, err := bcrypt.GeneratePassword(password)
	if err != nil {
		return err
	}
	return u.db.Update(&model.User{
		Model: base.Model{
			ID: userId,
		},
		Password: hashPwd,
	})
}

// 删除用户
func (u *UserService) Delete(id uint) error {
	return u.db.Delete(id)
}

// 设置头像
func (u *UserService) UploadAvatar(c *gin.Context, uid uint) error {
	file, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	// 存图片
	td := time.Now()
	fileName := md5.GenerateLower(strconv.Itoa(int(td.UnixMilli())) + strconv.Itoa(int(rand.Int())))
	path := strings.Builder{}
	path.WriteString("user/")
	path.WriteString(strconv.Itoa(td.Year()))
	path.WriteString("/")
	path.WriteString(strconv.Itoa(int(td.Month())))
	path.WriteString("/")
	path.WriteString(strconv.Itoa(td.Day()))
	path.WriteString("/")
	path.WriteString(fileName)
	path.WriteString(filepath.Ext(file.Filename))
	err = c.SaveUploadedFile(file, pkg.GetRootPath()+"/web/upload/"+path.String())
	if err != nil {
		logger.LogError(err)
		return err
	}

	// 存数据库
	user, err := u.db.GetById(uid)
	if err != nil {
		return err
	}
	avatar := path.String()
	user.Avatar = &avatar
	err = u.db.Update(user)
	if err != nil {
		return err
	}

	return nil
}
