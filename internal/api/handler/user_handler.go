package handler

import (
	"void-project/internal/api/request"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/model"
	"void-project/internal/service"
	"void-project/pkg/bcrypt"
	"void-project/pkg/jwt"
	"void-project/pkg/logger"

	"github.com/gin-gonic/gin"
)

type User struct {
	service *service.UserService
}

func NewUser() *User {
	return &User{service: service.NewUserService()}
}

// 注册
func (u *User) Register(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}
	if user.Account == "" || user.Password == "" || user.Identity == nil {
		response.FailError(c, apierr.MissingAcctPwd)
		return
	}
	if user.Password != *user.Identity {
		response.FailError(c, apierr.PasswordMismatch)
		return
	}

	if existsAccount := u.service.ExistsAccount(user.Account); existsAccount {
		response.FailError(c, apierr.AccountExists)
		return
	}

	err = u.service.Register(&user)
	if err != nil {
		response.FailError(c, apierr.RegisterFailed, err)
		return
	}
	response.Success(c, user)
}

// 登录
func (u *User) Login(c *gin.Context) {
	var param struct {
		Account  string
		Password string
	}
	if err := c.ShouldBind(&param); err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}
	if param.Account == "" || param.Password == "" {
		response.FailError(c, apierr.MissingAcctPwd)
		return
	}
	eu, err := u.service.GetByAccount(param.Account)
	if err != nil {
		response.FailError(c, apierr.LoginFailed, err)
		return
	}
	if eu == nil || eu.Account == "" {
		response.FailError(c, apierr.AccountNotExist)
		return
	}

	// ok := md5.CheckPassword(param.Password, *eu.Salt, eu.Password)
	if ok := bcrypt.ComparePassword(eu.Password, param.Password); !ok {
		response.FailError(c, apierr.InvalidPassword)
		return
	}

	user, err := u.service.GetByAccountPassword(param.Account, eu.Password)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.LoginFailed, err)
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InternalServerError, err)
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InternalServerError, err)
		return
	}

	response.Success(c, map[string]any{
		"token":       token,
		"user_id":     user.ID,
		"expire_time": claims.ExpiresAt.UnixMilli(),
	})
}

// 查询用户
func (u *User) Fetch(c *gin.Context) {
	id, err := request.GetParamIntErr(c, "id")
	if err != nil {
		return
	}
	user, err := u.service.Fetch(uint(id))
	if err != nil {
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.Success(c, user)
}

// 用户列表
func (u *User) List(c *gin.Context) {
	pager, err := request.PageQuery(c)
	if err != nil {
		return
	}
	users, total, err := u.service.List(pager)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.SuccessPage(c, users, total)
}

// 更新
func (u *User) Update(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBind(user)
	if err != nil {
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}
	user.ID = request.GetAuthUserId(c)

	err = u.service.Update(user)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.UpdateFailed, err)
		return
	}

	response.Success(c, user)
}

// 修改密码
func (u *User) UpdatePassword(c *gin.Context) {
	var param struct {
		Password string
	}
	if err := c.ShouldBind(&param); err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InvalidParameter, err)
		return
	}
	if param.Password == "" {
		response.FailError(c, apierr.MissingPwd)
		return
	}

	if err := u.service.UpdatePassword(request.GetAuthUserId(c), param.Password); err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.UpdateFailed, err)
		return
	}
	response.SuccessOk(c)
}

// 删除
func (u *User) Delete(c *gin.Context) {
	err := u.service.Delete(request.GetAuthUserId(c))
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.DeleteFailed, err)
		return
	}
	response.SuccessOk(c)
}

// 设置头像
func (u *User) Avatar(c *gin.Context) {
	err := u.service.UploadAvatar(c, request.GetAuthUserId(c))
	if err != nil {
		response.FailError(c, apierr.FileUploadFailed, err)
		return
	}
	response.SuccessOk(c)
}
