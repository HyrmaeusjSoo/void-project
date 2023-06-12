package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewHandler struct{}

func (*ViewHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func (*ViewHandler) Chat(c *gin.Context) {
	c.HTML(http.StatusOK, "chat.tmpl", nil)
}
