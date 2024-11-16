protoc --proto_path=. \
    --go_out=. --go_opt=Mbeef/beef.proto=./beef \
    --go-grpc_out=. --go-grpc_opt=Mbeef/beef.proto=./beef \
    beef/beef.proto