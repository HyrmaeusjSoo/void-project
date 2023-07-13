package handler

import (
	"strconv"

	"void-project/internal/api/request"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/model"
	"void-project/internal/service"
	"void-project/pkg/jwt"
	"void-project/pkg/logger"
	"void-project/pkg/md5"

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
		response.FailError(c, apierr.InvalidParameter)
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
		logger.LogError(err)
		response.FailError(c, apierr.CreateFailed)
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
		response.FailError(c, apierr.InvalidParameter)
		return
	}
	if param.Account == "" || param.Password == "" {
		response.FailError(c, apierr.MissingAccountPassword)
		return
	}
	data, err := u.service.GetByAccount(param.Account)
	if err != nil {
		response.FailError(c, apierr.InternalServerError)
		return
	}
	if data.Account == "" {
		response.FailError(c, apierr.AccountNotExist)
		return
	}

	ok := md5.CheckPassword(param.Password, *data.Salt, data.Password)
	if !ok {
		response.FailError(c, apierr.InvalidPassword)
		return
	}

	user, err := u.service.GetByAccountPassword(param.Account, data.Password)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InternalServerError)
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InternalServerError)
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InternalServerError)
		return
	}

	response.Success(c, map[string]any{
		"token":      token,
		"userId":     user.ID,
		"expireTime": claims.ExpiresAt.UnixMilli(),
	})
}

// 获取
func (u *User) Fetch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InvalidParameter)
		return
	}
	user, err := u.service.Fetch(uint(id))
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.FetchFailed)
		return
	}
	response.Success(c, user)
}

// 获取列表
func (u *User) List(c *gin.Context) {
	pager, err := request.PageQuery(c)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.InvalidParameter)
		return
	}
	users, total, err := u.service.List(*pager)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.FetchFailed)
		return
	}
	response.SuccessPage(c, users, total)
}

// 更新
func (u *User) Update(c *gin.Context) {
	id, err := request.GetParamIntErr(c, "id")
	if err != nil {
		return
	}

	user := &model.User{}
	err = c.ShouldBind(user)
	if err != nil {
		response.FailError(c, apierr.InvalidParameter)
		return
	}
	user.ID = uint(id)

	err = u.service.Update(user)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.UpdateFailed)
		return
	}

	response.Success(c, user)
}

// 删除
func (u *User) Delete(c *gin.Context) {
	id, err := request.GetParamIntErr(c, "id")
	if err != nil {
		return
	}

	err = u.service.Delete(uint(id))
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.DeleteFailed)
		return
	}
	response.SuccessOk(c)
}
