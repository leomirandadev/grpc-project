# GRPC Golang to Node.js
In this repository you will find two applications communicating with each other over gRPC.

On MacOs, install protobuf.
``` shell
$ brew install protobuf
```

## Node Server
Dependencies:
- @grpc/grpc-js ;
- @grpc/proto-loader ;


## Golang Client
Dependencies:
- google.golang.org/grpc ;

To create golang protoFile you need install first protoc-gen-go:
``` shell
$ brew install protoc-gen-go
```

After install protoc-gen-go, run on terminal:
```shell
$ protoc --go_out=plugins=grpc:posts_service posts_service.proto
```