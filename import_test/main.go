package main

import (
	"time"

	"github.com/IanVzs/virtual_room"
)

func main() {
	hub := virtual_room.NewHub()
	go hub.Run()

	virtual_room.ServeEvent(hub, "e1", "1")
	virtual_room.ServeEvent(hub, "e1", "1_2")
	virtual_room.ServeEvent(hub, "e2", "2")
	virtual_room.ServeEvent(hub, "e3", "3")

	time.Sleep(1 * time.Second)
}
