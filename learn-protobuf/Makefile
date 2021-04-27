gen:
  protoc --go_out=pb --go-grpc_out=pb --proto_path=proto proto/*.proto

clean:
  rm pb/*.go

test:
  go test -cover -race ./..

client:
  go run cmd/client/main.go -address 0.0.0.0:8888

server:
  go run cmd/server/main.go -port 8888