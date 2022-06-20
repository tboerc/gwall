protos:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
		protos/*.proto

server:
	$(MAKE) start -C server

client:
	$(MAKE) start -C client
