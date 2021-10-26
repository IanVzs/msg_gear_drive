package virtual_room

import (
	"testing"
	"time"
)

func TestAddWhileRunning(t *testing.T) {
	i := 1
	i2 := 2
	c := NewCron()
	c.AddFunc("@every 2s", func() {
		t.Log("每2秒执行一次", i2)
		i2 += 2
	})

	c.Start()

	time.Sleep(time.Second * 3)
	c.AddFunc("@every 1s", func() {
		t.Log("每秒执行一次", i)
		i++
	})
	time.Sleep(time.Second * 2)
}

func TestGCron(t *testing.T) {
	i := 1
	i_m := 1
	c := NewCron()
	c.AddFunc("@every 1s", func() {
		t.Log("每秒执行一次", i)
		i++
	})
	c.AddFunc("*/1 * * * *", func() {
		t.Log("每分执行一次 绝对秒 跟随系统时间", i_m)
		i_m++
	})
	c.AddFunc("35 17 * * *", func() {
		t.Log("定点执行", time.Now())
		i_m++
	})
	c.Start()
	// time.Sleep(time.Minute * 1)
	time.Sleep(time.Second * 1)
}
