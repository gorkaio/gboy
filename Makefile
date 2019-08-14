build:
	go build -o gboy cmd/gboy/main.go

test:
	go test -v pkg/cart/*.go
	go test -v pkg/memory/*.go
	go test -v pkg/cpu/*.go