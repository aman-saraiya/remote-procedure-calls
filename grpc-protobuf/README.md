go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

Commands to compile the proto files

protoc --go_out=../ --go-grpc_out=../ statsservice.proto

protoc --go_out=../ --go-grpc_out=../ userservice.proto
