all: test
test:
	go test vroom_test.go hub.go client.go vroom.go

test_v:
	go test -v vroom_test.go hub.go client.go vroom.go

test_i:
	mv test_import ..
	cd ../test_import && go run main.go
	mv ../test_import .

help:
	@echo "make: test all"
	@echo "make test_v: test with log"
	@echo "make test_i: test import this mod"