start:
	go run .

build:
	go build .

test:
	go test -v ./...

coverage:
	go test -coverprofile='coverage.out' ./...
	go tool cover -html='coverage.out'