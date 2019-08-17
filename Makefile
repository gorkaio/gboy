build: generate
	go build -o gboy cmd/gboy/main.go

test: generate
	go test ./...

generate:
	go generate ./...