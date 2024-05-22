build:
	go build main.go

check:
	gopls check main.go

fmt:
	go fmt main.go

test:
	go test
