package handler

import (
	"net/http"
	"strconv"

	"void-project/internal/api/request"
	"void-project/internal/api/response"
	"void-project/internal/model"
	"void-project/internal/service"
	"void-project/pkg/jwt"
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
		response.Fail(c, http.StatusOK, "参数有误！")
	}
	if user.Account == "" || user.Password == "" || user.Identity == nil {
		response.Fail(c, http.StatusOK, "账号或密码不能为空！")
		return
	}
	if user.Password != *user.Identity {
		response.Fail(c, http.StatusOK, "两次密码不一致！")
		return
	}

	if existsAccount := u.service.ExistsAccount(user.Account); existsAccount {
		response.Fail(c, http.StatusOK, "用户已存在！")
		return
	}

	err = u.service.Register(&user)
	if err != nil {
		response.Fail(c, http.StatusOK, "保存失败"+err.Error())
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
		response.Fail(c, http.StatusForbidden, err.Error())
		return
	}
	if param.Account == "" || param.Password == "" {
		response.Fail(c, http.StatusForbidden, "账号和密码不能为空！")
		return
	}
	data, err := u.service.GetByAccount(param.Account)
	if err != nil {
		response.Fail(c, http.StatusOK, "登录失败:"+err.Error())
		return
	}
	if data.Account == "" {
		response.Fail(c, http.StatusOK, "账号不存在")
		return
	}

	ok := md5.CheckPassword(param.Password, *data.Salt, data.Password)
	if !ok {
		response.Fail(c, http.StatusOK, "密码错误")
		return
	}

	user, err := u.service.GetByAccountPassword(param.Account, data.Password)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}

	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	claims, err := jwt.ParseToken(token)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
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
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := u.service.Fetch(uint(id))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, user)
}

// 获取列表
func (u *User) List(c *gin.Context) {
	pager, err := request.PageQuery(c)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	users, total, err := u.service.List(*pager)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.SuccessPage(c, users, total)
}

// 更新
func (u *User) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	user := &model.User{}
	err = c.ShouldBind(user)
	if err != nil {
		response.Fail(c, http.StatusBadRequest, "参数有误！")
		return
	}
	user.ID = uint(id)

	err = u.service.Update(user)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}

	response.Success(c, user)
}

// 删除
func (u *User) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = u.service.Delete(uint(id))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, "删除成功")
}
