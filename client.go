package virtual_room

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron/v3"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 3 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type todoNote struct {
	Status    int // 未进行/进行中/已完成/已放弃/已失败
	Name      string
	T         string    // 类型
	Condition [4]string // 达成条件 时间/地点/人物/事件
}
type handNote struct {
	stack *Stack
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	sync.Mutex
	hub *Hub
	// crontab
	crontab *cron.Cron
	// ...
	handNote *handNote

	// Buffered channel of outbound messages.
	receive chan []byte
	beat    chan struct{}
	info    *Object // 只可读,不可写
	name    string
}

// func (c *Client) setHandNote(mem []byte) {
// 	c.Lock()
// 	defer c.Unlock()
// }

func (c *Client) heartbeat() {
	ticker := time.NewTicker(pingPeriod)
	rSign := false
	defer func() {
		ticker.Stop()
		c.hub.unregister <- c
	}()
	for {
		bfselect := time.Now().UnixNano()
		select {
		case _, ok := <-c.beat:
			if !ok {
				fmt.Printf("Channel 错误关闭(%s): %d\n", c.name, time.Now().UnixNano())
				return
			}
			rSign = true
			fmt.Printf("心跳(%s) - 时间Unix时间: %d\n", c.name, time.Now().UnixNano())
			fmt.Printf("基本信息更新: %+v\n", *c.info)
		case <-ticker.C:
			if rSign {
				rSign = false
				ticker.Reset(pingPeriod)
				continue
			}
			_d := time.Now().UnixNano()
			cose := _d - bfselect
			fmt.Printf("心脏骤停(%s) - Cose: %d | Now: %d\n", c.name, cose, _d)
			return
		}
	}
}

func (c *Client) writePump() {
	defer func() {
		c.hub.unregister <- c
	}()
	for {
		message, ok := <-c.receive
		if !ok {
			// The hub closed the channel.
			fmt.Println("心脏骤停,writePump退出")
			return
		}
		fmt.Printf("%s -收到- %s - 时间Unix时间: %d\n", c.name, message, time.Now().UnixNano())
	}
}

// TODO 写入csv中定时事件/有bug
func (c *Client) writeCronTab() {
	eventName := "TODO测试事件"
	c.crontab.AddFunc("@every 2s", func() {
		peek := c.handNote.stack.Peek()
		if _peek, ok := peek.(todoNote); ok {
			peek_str, _ := json.Marshal(_peek)
			// fmt.Println(_peek, peek_str)
			c.hub.broadcast <- []byte(peek_str)
		}
	})
	c.handNote.stack.Push(todoNote{Status: 0, Name: c.name + eventName, T: "test", Condition: [4]string{"*", "*", "test", "happend"}})
}
