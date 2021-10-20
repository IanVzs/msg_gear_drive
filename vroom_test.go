package virtual_room

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestVRoom(t *testing.T) {
	hub := NewHub()
	go hub.run()

	// ServeEvent(hub, "e1", "1")
	// ServeEvent(hub, "e1", "1_2")
	// ServeEvent(hub, "e2", "2")
	// ServeEvent(hub, "e3", "3")

	for _, i := range []int{1, 2, 3, 4, 5, 6, 7} {
		s_i := strconv.Itoa(i)
		ServeEvent(hub, "e"+s_i, s_i)
		time.Sleep(30 * time.Millisecond)
	}
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------------")
	for _, i := range []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 2, 3} {
		s_i := strconv.Itoa(i)
		ServeEvent(hub, "e"+s_i, s_i)
		time.Sleep(30 * time.Millisecond)
	}

	time.Sleep(2 * time.Second)
	t.Log("Test Done")
	// time.Sleep(time.Second * 1)
}
