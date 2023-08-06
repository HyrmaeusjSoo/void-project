package handler

import (
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

// 发送消息
func (m *Message) SendUserMsg(c *gin.Context) {
	err := m.service.Chat(c.Writer, c.Request)
	if err != nil {
		logger.LogError(err)
	}
}

// 查询在线用户
func (m *Message) OnLine(c *gin.Context) {
	id := request.GetAuthUserId(c)

	users, err := m.service.OnLine(id)
	if err != nil {
		response.FailError(c, apierr.FetchFailed, err)
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
		response.FailError(c, apierr.FetchFailed, err)
		return
	}
	response.SuccessCursor(c, messages, next)
}
