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
	pongWait = 1 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type todoNote struct {
	status    int // 未进行/进行中/已完成/已放弃/已失败
	name      string
	t         string    // 类型
	condition [4]string // 达成条件 时间/地点/人物/事件
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
	name    string
}

func (c *Client) setHandNote(mem []byte) {
	c.Lock()
	defer c.Unlock()
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	rSign := false
	defer func() {
		ticker.Stop()
		c.hub.unregister <- c
	}()
	for {
		bfselect := time.Now().UnixNano()
		select {
		case message, ok := <-c.receive:
			if !ok {
				// The hub closed the channel.
				c.setHandNote([]byte("hub 关闭"))
				return
			}
			rSign = true
			fmt.Printf("%s -收到- %s - 时间Unix时间: %d\n", c.name, message, time.Now().UnixNano())
		case <-ticker.C:
			if rSign {
				rSign = false
				ticker.Reset(pingPeriod)
				continue
			}
			_d := time.Now().UnixNano()
			cose := _d - bfselect
			fmt.Printf("%s -- %s , Cose: %d | Now: %d\n", c.name, []byte("Time 关闭"), cose, _d)
			return
		}
	}
}

// TODO 写入csv中定时事件/有bug
func (c *Client) writeCronTab() {
	eventName := "TODO测试事件"
	c.crontab.AddFunc("@every 2s", func() {
		peek := c.handNote.stack.Peek()
		peek_str, _ := json.Marshal(peek)
		fmt.Println(peek, peek_str)
		c.hub.broadcast <- []byte(peek_str)
	})
	c.handNote.stack.Push(todoNote{status: 0, name: eventName, t: "test", condition: [4]string{"*", "*", "test", "happend"}})
}
