package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"chat/internal/api/response"
	"chat/internal/model"
	"chat/internal/service"
	"chat/pkg"

	"github.com/gin-gonic/gin"
)

type User struct{}

var userService = &service.UserService{}

// 注册
func (*User) Register(ctx *gin.Context) {
	account := ctx.Request.FormValue("account")
	password := ctx.Request.FormValue("password")
	repassword := ctx.Request.FormValue("identity")
	name := ctx.Request.FormValue("name")
	if account == "" || password == "" || repassword == "" {
		response.Fail(ctx, http.StatusOK, "账号或密码不能为空！")
		return
	}
	if password != repassword {
		response.Fail(ctx, http.StatusOK, "两次密码不一致！")
		return
	}

	if existsAccount := userService.ExistsAccount(account); existsAccount {
		response.Fail(ctx, http.StatusOK, "用户已存在！")
		return
	}

	salt := fmt.Sprintf("%d", rand.Int31())
	t := time.Now()
	user := model.User{
		Account:       account,
		Password:      pkg.SaltPassword(password, salt),
		Salt:          salt,
		Name:          name,
		LoginTime:     &t,
		LoginOutTime:  &t,
		HeartBeatTime: &t,
	}
	userService.Register(&user)
	response.Success(ctx, user)
}

// 登录
func (*User) Login(ctx *gin.Context) {
	var param struct {
		Account  string
		Password string
	}
	if err := ctx.ShouldBind(&param); err != nil {
		response.Fail(ctx, http.StatusForbidden, err.Error())
		return
	}
	if param.Account == "" || param.Password == "" {
		response.Fail(ctx, http.StatusForbidden, "账号和密码不能为空！")
		return
	}
	data, err := userService.GetByAccount(param.Account)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "登录失败:" + err.Error(),
		})
		return
	}
	if data.Account == "" {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "账号不存在",
		})
		return
	}

	/* ok := pkg.CheckPassword(param.Password, data.Salt, data.Password)
	if !ok {
		ctx.JSON(200, gin.H{
			"code":    -1,
			"message": "密码错误",
		})
		return
	}

	rsp, err := repo.FindUserByAccountAndPwd(param.Account, data.Password)
	if err != nil {
		zap.S().Info("登录失败", err)
	}

	token, err := middleware.GenerateToken(rsp.ID, "yk")
	if err != nil {
		zap.S().Info("生成token失败", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "登录成功",
		"tokens":  token,
		"userId":  rsp.ID,
	}) */
}

// 获取
func (*User) Fetch(ctx *gin.Context) {
	user, err := userService.Fetch(1)
	if err != nil {
		response.Fail(ctx, 10003, err.Error())
		return
	}
	response.Success(ctx, user)
}

// 下载文件
func (*User) Download(ctx *gin.Context) {
	ctx.File("D:/abc.xlsx")
}
