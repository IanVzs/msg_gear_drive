// 2021-10-21: go get github.com/robfig/cron/v3@v3.0.0
package virtual_room

import "github.com/robfig/cron/v3"

// var secondParser = cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)

func NewCron() *cron.Cron {
	return cron.New()
}
