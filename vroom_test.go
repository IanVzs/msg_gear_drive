package virtual_room

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestVRoom(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	info := Object{ID: "1", Name: "e1"}
	ServeEvent(hub, info.Name, info.ID, info)
	info.ID = "1_2"
	ServeEvent(hub, info.Name, info.ID, info)
	info.ID = "2"
	info.Name = "e" + info.ID
	ServeEvent(hub, info.Name, info.ID, info)
	info.ID = "3"
	info.Name = "e" + info.ID
	ServeEvent(hub, info.Name, info.ID, info)
	// ServeEvent(hub, "e1", "1_2")
	// ServeEvent(hub, "e2", "2")
	// ServeEvent(hub, "e3", "3")

	time.Sleep(5 * time.Second)
	t.Log("Test Done")
	// time.Sleep(time.Second * 1)
}

func TestVRoomByLoop(t *testing.T) {
	hub := NewHub()
	go hub.Run()

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		s_i := strconv.Itoa(i)
		info := Object{ID: s_i, Name: "e" + s_i}
		ServeEvent(hub, info.Name, info.ID, info)
		// ServeEvent(hub, "e"+s_i, s_i)

		time.Sleep(300 * time.Millisecond)
	}
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------")
	for _, i := range []int{1, 2, 3, 1, 2, 3} {
		s_i := strconv.Itoa(i)
		info := Object{ID: s_i, Name: "e" + s_i}
		ServeEvent(hub, info.Name, info.ID, info)
		// ServeEvent(hub, "e"+s_i, s_i)
		time.Sleep(300 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)
	t.Log("Test Done")
}
