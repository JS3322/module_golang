/*
brew install protobuf


go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

/*

system must have following installed
- protobuf
- protoc-gen-go
- protoc-gen-go-grpc

protoc --go-grpc_out=require_unimplemented_servers=false:./chatserver/ --go_out=./chatserver/ chat.proto