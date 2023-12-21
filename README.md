# Protobuf ServiceStatus

```shell
protoc --proto_path=. \
  --go_out=./service_status --go_opt=paths=source_relative \
  --go-grpc_out=./service_status --go-grpc_opt=paths=source_relative \
  service_status.proto

```