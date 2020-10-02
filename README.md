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

To install protofile you need install first protoc-gen-go, like:
``` shell
$ brew install protoc-gen-go
```

To install protofile and create client, run on terminal
```shell
$ protoc -I . posts_service.proto --go_out=plugins=grpc:./go-client/services
```