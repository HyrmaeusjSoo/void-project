package handler

import (
	"net/http"
	"strconv"

	"chat/internal/api/response"
	"chat/internal/model"
	"chat/internal/service"
	"chat/pkg/jwt"
	"chat/pkg/md5"

	"github.com/gin-gonic/gin"
)

type User struct{}

var userService = &service.UserService{}

// 注册
func (*User) Register(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBind(&user)
	if err != nil {
		response.Fail(c, http.StatusOK, "参数有误！")
	}
	if user.Account == "" || user.Password == "" || user.Identity == "" {
		response.Fail(c, http.StatusOK, "账号或密码不能为空！")
		return
	}
	if user.Password != user.Identity {
		response.Fail(c, http.StatusOK, "两次密码不一致！")
		return
	}

	if existsAccount := userService.ExistsAccount(user.Account); existsAccount {
		response.Fail(c, http.StatusOK, "用户已存在！")
		return
	}

	err = userService.Register(&user)
	if err != nil {
		response.Fail(c, http.StatusOK, "保存失败"+err.Error())
		return
	}
	response.Success(c, user)
}

// 登录
func (*User) Login(c *gin.Context) {
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
	data, err := userService.GetByAccount(param.Account)
	if err != nil {
		response.Fail(c, http.StatusOK, "登录失败:"+err.Error())
		return
	}
	if data.Account == "" {
		response.Fail(c, http.StatusOK, "账号不存在")
		return
	}

	ok := md5.CheckPassword(param.Password, data.Salt, data.Password)
	if !ok {
		response.Fail(c, http.StatusOK, "密码错误")
		return
	}

	user, err := userService.GetByAccountPassword(param.Account, data.Password)
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
func (*User) Fetch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}
	user, err := userService.Fetch(uint(id))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, user)
}

// 获取列表
func (*User) List(c *gin.Context) {
	users, err := userService.List()
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, users)
}

// 更新
func (*User) Update(c *gin.Context) {
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

	err = userService.Update(user)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}

	response.Success(c, user)
}

// 删除
func (*User) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = userService.Delete(uint(id))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, "删除成功")
}
