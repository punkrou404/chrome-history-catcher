
all: build deploy
build:
	go get github.com/mattn/go-sqlite3
	go build src/main.go
	mv main chc
deploy:
	sudo cp chc /usr/local/bin/