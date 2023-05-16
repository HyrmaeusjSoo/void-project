package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	account := c.Request.FormValue("account")
	password := c.Request.FormValue("password")
	repassword := c.Request.FormValue("identity")
	name := c.Request.FormValue("name")
	if account == "" || password == "" || repassword == "" {
		response.Fail(c, http.StatusOK, "账号或密码不能为空！")
		return
	}
	if password != repassword {
		response.Fail(c, http.StatusOK, "两次密码不一致！")
		return
	}

	if existsAccount := userService.ExistsAccount(account); existsAccount {
		response.Fail(c, http.StatusOK, "用户已存在！")
		return
	}

	salt := fmt.Sprintf("%d", rand.Int31())
	t := time.Now()
	user := model.User{
		Account:       account,
		Password:      md5.SaltPassword(password, salt),
		Salt:          salt,
		Name:          name,
		LoginTime:     &t,
		LoginOutTime:  &t,
		HeartBeatTime: &t,
	}
	userService.Register(&user)
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
		response.Fail(c, http.StatusOK, err.Error())
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

	response.Success(c, map[string]any{
		"token":  token,
		"userId": user.ID,
	})
}

// 获取
func (*User) Fetch(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	user, err := userService.Fetch(uint(id))
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, user)
}

// 更新
func (*User) Update(c *gin.Context) {
	c.Request.FormValue("")
	response.Success(c, nil)
}
