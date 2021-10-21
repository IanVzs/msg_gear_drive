all: tvroom
tvroom:
	go test vroom_test.go hub.go client.go vroom.go gcron.go
tcron:
	go test -v cron_test.go gcron.go
tcron_addwhilerunning:
	go test -v cron_test.go gcron.go -run TestAddWhileRunning

test_v:
	go test -v vroom_test.go hub.go client.go vroom.go gcron.go

test_i:
	mv import_test ..
	cd ../import_test && go run main.go
	mv ../import_test .

help:
	@echo "make: test all"
	@echo "make tvroom: test vroom"
	@echo "make tcron: test crontab"
	@echo "make test_i: test import this mod"