protoc --proto_path=. \
    --go_out=paths=source_relative:. \
    --go-http_out=paths=source_relative:. \
    --go-grpc_out=paths=source_relative:. \
    ./user.proto