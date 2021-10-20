package virtual_room

import (
	"fmt"
	"sync"
	"time"
)

const (
	// Time allowed to read the next pong message from the peer.
	pongWait = 100 * time.Millisecond

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
)

type handNote struct {
	status int
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	sync.Mutex
	hub *Hub
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
		c.handNote.status = 0
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
