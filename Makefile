build: generate
	go build -o gboy cmd/gboy/main.go

test:
	go test -coverprofile=coverage.out ./...

generate:
	go generate ./...