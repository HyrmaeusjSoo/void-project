package handler

import (
	"fmt"
	"net/http"
	"void-project/internal/api/request"
	"void-project/internal/api/response"
	"void-project/internal/api/response/apierr"
	"void-project/internal/service"
	"void-project/pkg/logger"

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

// 消息列表
func (m *Message) List(c *gin.Context) {
	cursor := request.CursorQuery(c)
	uId := request.GetAuthUserId(c)
	targetId := request.GetQueryInt(c, "target_id")
	messages, next, err := m.service.List(uId, uint(targetId), cursor)
	if err != nil {
		logger.LogError(err)
		response.FailError(c, apierr.FetchFailed)
		return
	}
	response.SuccessCursor(c, messages, next)
}
