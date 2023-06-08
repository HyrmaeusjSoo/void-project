package handler

import (
	"chat/internal/service"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

type Message struct{}

var msgService service.MessageService

func (*Message) SendUserMsg(c *gin.Context) {
	/* err := msgService.Chat(c.Writer, c.Request)
	if err != nil {
		fmt.Println(err.Error())
	} */
	chat(c.Writer, c.Request)
}

func chat(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		panic(err)
	}
	defer conn.Close(websocket.StatusInternalError, "websocket已关闭")

	var sr = func(reqCtx context.Context) {
		ctx, cancel := context.WithTimeout(reqCtx, time.Minute*30)
		defer cancel()
		//接收信息
		typ, data, err := conn.Read(ctx)
		if err != nil {
			panic(err)
		}
		//返回信息
		str := string(data) + ", resMsg:回复"
		if err := conn.Write(ctx, typ, []byte(str)); err != nil {
			panic(err)
		}
	}

	for {
		sr(r.Context())
	}

}
