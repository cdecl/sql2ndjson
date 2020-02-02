
PROJECT=sql2ndjson
BIN=$(CURDIR)/bin
EXEC=$(PROJECT)


all: build 

build:
	go build -o $(BIN)/$(EXEC).exe

test:
	go test -v 

dep:
	go mod tidy
	
cc:
	SET GOOS=linux& SET GOARCH=amd64& go build -o $(BIN)/$(EXEC) 
	SET GOOS=windows& SET GOARCH=amd64& go build -o $(BIN)/$(EXEC).exe