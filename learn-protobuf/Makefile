gen:
	protoc --go_out=pb --go-grpc_out=pb --proto_path=proto proto/*.proto

mv:
	mv pb/proto/*.go pb/
clean:
	rm pb/*.go

test:
	go test -cover -race ./..

client:
	go run cmd/client/main.go -address 127.0.0.1:8888

server:
	go run cmd/server/main.go -port 8888
