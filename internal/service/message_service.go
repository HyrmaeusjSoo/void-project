package service

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"void-project/internal/model"
	"void-project/internal/model/base"
	"void-project/internal/repository/mysql"
	"void-project/pkg/logger"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type MessageService struct {
	db  *mysql.MessageRepository
	udb *mysql.UserRepository
}

func NewMessageService() *MessageService {
	return &MessageService{mysql.NewMessageRepository(), mysql.NewUserRepository()}
}

type Node struct {
	Conn     *websocket.Conn
	Addr     string
	MsgQueue chan *model.Message //真消息队列
}

var (
	clientMap = make(map[uint]*Node, 10) //客户端map
	rwLocker  sync.RWMutex               //线程安全读写锁
)

// 聊天
func (m *MessageService) Chat(w http.ResponseWriter, r *http.Request) error {
	IdStr := r.URL.Query().Get("user_id")
	if IdStr == "" {
		return errors.New("user_id为空")
	}
	id, err := strconv.ParseUint(IdStr, 10, 32)
	if err != nil {
		return errors.New("user_id类型转换失败" + err.Error())
	}
	userId := uint(id)

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{InsecureSkipVerify: true})
	if err != nil {
		return errors.New("升级连接到websocket失败" + err.Error())
	}
	fmt.Println("WebSocket开启", userId, "----------------------")

	node := &Node{
		Conn:     conn,
		Addr:     r.RemoteAddr,
		MsgQueue: make(chan *model.Message, 32),
	}
	rwLocker.Lock()
	clientMap[userId] = node
	rwLocker.Unlock()

	go m.send(r.Context(), node)
	err = m.receive(r.Context(), node)

	fmt.Println("WebSocket关闭", userId, "=====================")
	if err != nil {
		conn.Close(websocket.StatusInternalError, "接收信息错误"+err.Error())
	}
	conn.Close(websocket.StatusNormalClosure, "websocket已关闭")

	return err
}

// 发
func (m *MessageService) send(ctx context.Context, node *Node) {
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

// 收
func (m *MessageService) receive(ctx context.Context, node *Node) error {
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
			return errors.New("对方已离开" + strconv.Itoa(int(msg.TargetId)))
		}
		tarNode.MsgQueue <- &msg

		err = m.db.Create(&msg)
		if err != nil {
			logger.LogError(err)
		}
	}
}

// 在线用户
func (m *MessageService) OnLine(userId uint) ([]model.User, error) {
	ids := make([]uint, 0, len(clientMap))
	for k := range clientMap {
		if k == userId {
			continue
		}
		ids = append(ids, k)
	}
	if len(ids) == 0 {
		return nil, nil
	}
	return m.udb.GetInIds(ids)
}

// 消息列表
func (m *MessageService) List(uId, targetId uint, cursor base.Cursor) ([]model.Message, base.Next, error) {
	return m.db.GetListClean(uId, targetId, cursor)
}
