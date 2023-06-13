package service

import (
	"chat/internal/model"
	"chat/internal/repository/mysql"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type MessageService struct {
	db *mysql.UserRepository
}

func NewMessageService() *MessageService {
	return &MessageService{db: mysql.NewUserRepository()}
}

type Node struct {
	Conn     *websocket.Conn
	Addr     string
	MsgQueue chan *model.Message //真消息队列
}

var clientMap = make(map[uint]*Node, 10)
var rwLocker sync.RWMutex //线程安全读写锁

// 在线用户
func (m *MessageService) OnLine(selfid uint) ([]*model.User, error) {
	ids := make([]uint, 0, len(clientMap))
	for k := range clientMap {
		if k == selfid {
			continue
		}
		ids = append(ids, k)
	}
	if len(ids) == 0 {
		return nil, nil
	}
	return m.db.GetInIds(ids)
}

// 聊天
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
	fmt.Println("开启----------------------")

	node := &Node{
		Conn:     conn,
		Addr:     r.RemoteAddr,
		MsgQueue: make(chan *model.Message, 32),
	}
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	go send(r.Context(), node)
	err = receive(r.Context(), node)

	fmt.Println("关闭=====================")
	if err != nil {
		conn.Close(websocket.StatusInternalError, "接收信息错误"+err.Error())
	}
	conn.Close(websocket.StatusNormalClosure, "websocket已关闭")

	return err
}

func send(ctx context.Context, node *Node) {
	/* for {
		select {
		case data := <-node.MsgQueue:
			// err := node.Conn.Write(ctx, websocket.MessageText, data)
			err := wsjson.Write(ctx, node.Conn, data)
			if err != nil {
				return
			}
		}
	} */
	for msg := range node.MsgQueue {
		err := wsjson.Write(ctx, node.Conn, msg)
		if err != nil {
			return
		}
	}
}

func receive(ctx context.Context, node *Node) error {
	for {
		msg := model.Message{}
		err := wsjson.Read(ctx, node.Conn, &msg)
		if err != nil {
			if errors.As(err, &websocket.CloseError{}) {
				return nil
			} else if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}
		tarNode, ok := clientMap[msg.TargetId]
		if !ok {
			return errors.New("对方已离开")
		}
		tarNode.MsgQueue <- &msg
	}
}
