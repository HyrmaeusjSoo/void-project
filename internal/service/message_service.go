package service

import (
	"chat/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"nhooyr.io/websocket"
)

type MessageService struct{}

type Node struct {
	Conn     *websocket.Conn
	Addr     string
	MsgQueue chan []byte //真消息队列
}

var clientMap = make(map[uint]*Node, 0)
var rwLocker sync.RWMutex //线程安全读写锁

func (*MessageService) Chat(w http.ResponseWriter, r *http.Request) error {
	IdStr := r.URL.Query().Get("userId")
	if IdStr == "" {
		return errors.New("userId为空")
	}
	id, err := strconv.ParseUint(IdStr, 10, 32)
	if err != nil {
		return errors.New("userId类型转换失败" + err.Error())
	}
	userId := uint(id)

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		return errors.New("升级连接到websocket失败" + err.Error())
	}
	// defer conn.Close(websocket.StatusInternalError, "websocket已关闭")

	node := &Node{
		Conn:     conn,
		Addr:     r.RemoteAddr,
		MsgQueue: make(chan []byte, 512),
	}
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	go sendProc(node, r.Context())
	go recProc(node, r.Context())

	return nil
}

func sendProc(node *Node, ctx context.Context) {
	for {
		fmt.Println("send------------")
		select {
		case data := <-node.MsgQueue:
			err := node.Conn.Write(ctx, websocket.MessageText, data)
			if err != nil {
				return
			}
		}
	}
}

func recProc(node *Node, ctx context.Context) {
	fmt.Println("rec--------------")
	for {
		// err := wsjson.Read(ctx, node.Conn, &msg)
		_, data, err := node.Conn.Read(ctx)
		if err != nil {
			return
		}
		msg := model.Message{}
		err = json.Unmarshal(data, &msg)
		if err != nil {
			return
		}
		fmt.Println("target:", msg.TargetId)
		tarNode, ok := clientMap[msg.TargetId]
		if !ok {
			return
		}
		tarNode.MsgQueue <- data
	}
}
