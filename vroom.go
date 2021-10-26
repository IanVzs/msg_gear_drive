package virtual_room

// ServeEvent handles Event from the peer.
func ServeEvent(hub *Hub, name string, event string, info Object) {
	for _c, _status := range hub.clients {
		if _c.name == name && _status {
			_c.beat <- struct{}{}
			_c.info = &info
			// hub.broadcast <- []byte(event)
			return
		}
	}
	client := &Client{hub: hub, handNote: &handNote{stack: NewStack()}, receive: make(chan []byte, 256), name: name, crontab: NewCron(), beat: make(chan struct{})}
	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.heartbeat()
	client.hub.register <- client
	// fmt.Printf("%s: 时间UnixNano时间: %d\n", event, time.Now().UnixNano())
	hub.broadcast <- []byte(event)
}
