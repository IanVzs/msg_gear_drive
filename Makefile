all: test
test:
	go test vroom_test.go hub.go client.go vroom.go

test_v:
	go test -v vroom_test.go hub.go client.go vroom.go
help:
	@echo "make: test all"