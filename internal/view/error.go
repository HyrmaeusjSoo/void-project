package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 未登录或404默认跳首页
func ErrorPage(c *gin.Context) {
	c.Redirect(http.StatusFound, "/web/index")
}
