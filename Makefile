
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u google.golang.org/grpc
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/gofiber/fiber/v2

.PHONY: server
server:
	go run main.go

.PHONY: server
client:
	go run client/client.go

.PHONY: proto
proto:
	protoc --proto_path=. --go_out=plugins=grpc:gen proto/user.proto
	
.PHONY: build
build:
	go build -o sample-grpc-go *.go