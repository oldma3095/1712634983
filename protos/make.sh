GOPATH:=$(shell go env GOPATH)
API_PROTO_FILES=$(shell find proto_files -name *.proto)

init:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

api:
protoc -I=./grpc/protos/third_party -I=./grpc/protos/ --go_out=./grpc/protos/ --go-grpc_out=./grpc/protos/ ./grpc/protos/common/*.proto