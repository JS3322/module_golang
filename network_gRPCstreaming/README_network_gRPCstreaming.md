system must have following installed
- protobuf
- protoc-gen-go
- protoc-gen-go-grpc

protoc --go-grpc-out=require_unimplemented_servers=false:./chatserver/ --go_out=./chatserver/ chat.proto