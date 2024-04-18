GOPATH:=$(shell go env GOPATH)
API_PROTO_FILES=$(shell find proto_files -name *.proto)

init:
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

api:
protoc -I=./protos/third_party -I=./protos/ --go_out=./protos/ --go-grpc_out=./protos/ ./protos/common/*.proto