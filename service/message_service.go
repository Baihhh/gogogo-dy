package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RaymondCode/simple-demo/models"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var messageIdSequence = int64(1)
var chatConnMap = sync.Map{}

type MessageResponse models.MessageResponse

func RunMessageServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:9090")
	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		var event = models.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := models.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}

func MessageSend(userId int64, toUserId int64, content string, actionType string) error {
	if actionType == "1" {
		atomic.AddInt64(&messageIdSequence, 1)
		curMessage := models.Message{
			//Id:         messageIdSequence,
			MsgContent: content,
			CreateTime: time.Now().Format("2006-01-02 15:04:05"),
			UserId:     userId,
			ToUserId:   toUserId,
		}
		err := send(curMessage)
		if err != nil {
			return err
		}
		return models.AddMessage(&curMessage)

	} else {
		return errors.New("不正确的消息类型")
	}
	return nil
}

func send(curMessage models.Message) error {
	//使用Dial函数和服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:9090")
	if err != nil {
		return err
	}
	//最后退出时关闭连接
	defer conn.Close()

	// 将结构体转换为字节切片
	jsonData, err := json.Marshal(curMessage)
	if err != nil {
		return err
	}
	//向服务端发送信息
	conn.Write(jsonData)
	return nil
}

func (q *MessageResponse) GetMessage(userId int64, toUserId int64) error {
	if userId != toUserId {
		return models.GetMessages(userId, toUserId, &q.MessageList)
	}
	return errors.New("userId 不能相同")
}
