package handler

import (
	"fmt"
	"net/http"
	"void-project/internal/api/request"
	"void-project/internal/api/response"
	"void-project/internal/service"

	"github.com/gin-gonic/gin"
)

type Message struct {
	service *service.MessageService
}

func NewMessage() *Message {
	return &Message{service.NewMessageService()}
}

func (m *Message) SendUserMsg(c *gin.Context) {
	err := m.service.Chat(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func (m *Message) OnLine(c *gin.Context) {
	id := request.GetAuthUserId(c)

	users, err := m.service.OnLine(id)
	if err != nil {
		response.Fail(c, http.StatusOK, err.Error())
		return
	}
	response.Success(c, users)

}
