DIST=.\dist
SERVER_FILE=$(DIST)\server.exe

protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		protos/*.proto

prepare:
	IF NOT EXIST "$(DIST)" MKDIR $(DIST)

clean-server: prepare
	IF EXIST "$(SERVER_FILE)" DEL $(SERVER_FILE)

build-server: clean-server
	go build -o $(SERVER_FILE) .\server

server: build-server
	$(SERVER_FILE)