build:
	go build main.go

check:
	gopls check main.go

format:
	go fmt main.go

test:
	go test
