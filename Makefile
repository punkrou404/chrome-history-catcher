all: test build deploy
build:
	go get -v github.com/mattn/go-sqlite3
	go build src/main/main.go
	mv main chc
deploy:
	sudo cp chc /usr/local/bin/
test:
	go test -v -cover ./...
test-cover:
	go test -coverprofile=c.out ./...
	go tool cover -func=c.out