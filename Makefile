DIST=.\dist
SERVER_FILE=$(DIST)\server.exe
CLIENT_FILE=$(DIST)\client.exe

.PHONY: proto server client
.SILENT: prepare clean-server build-server server clean-client build-client client

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		shared/proto/*.proto

prepare:
	IF NOT EXIST "$(DIST)" MKDIR $(DIST)

clean-server: prepare
	IF EXIST "$(SERVER_FILE)" DEL $(SERVER_FILE)

build-server: clean-server
	go build -o $(SERVER_FILE) .\server

server: build-server
	$(SERVER_FILE)

clean-client: prepare
	IF EXIST "$(CLIENT_FILE)" DEL $(CLIENT_FILE)

build-client: clean-client
	go build -o $(CLIENT_FILE) .\client

client: build-client
	$(CLIENT_FILE)