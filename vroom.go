package virtual_room

import (
	"fmt"
	"time"
)

// ServeEvent handles Event from the peer.
func ServeEvent(hub *Hub, name string, event string) {
	for _c, _status := range hub.clients {
		if _c.name == name && _status {
			hub.broadcast <- []byte(event)
			return
		}
	}
	client := &Client{hub: hub, handNote: &handNote{status: 1}, receive: make(chan []byte, 256), name: name}
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	client.hub.register <- client
	fmt.Printf("%s: 时间UnixNano时间: %d\n", event, time.Now().UnixNano())
	client.writeCronTab()
	client.hub.crontab.Start()
	hub.broadcast <- []byte(event)
}
