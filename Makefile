build: generate
	go build -o gboy cmd/gboy/main.go

test:
	go test ./...

generate:
	go generate ./...