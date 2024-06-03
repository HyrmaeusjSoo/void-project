package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewHandler struct{}

// 首页
func (*ViewHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

// WebSocket聊天示例
func (*ViewHandler) Chat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.tmpl", nil)
}

// API 请求地址和示例列表
func (*ViewHandler) Api(c *gin.Context) {
	c.HTML(http.StatusOK, "api.tmpl", nil)
}

// 存储页面
func (*ViewHandler) Storage(c *gin.Context) {
	c.HTML(http.StatusOK, "storage.tmpl", nil)
}
