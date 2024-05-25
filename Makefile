build:
	go build main.go broker.go

check:
	gopls check main.go broker.go

fmt:
	go fmt main.go broker.go broker_test.go

test:
	go test
