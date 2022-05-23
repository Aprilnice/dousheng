protoc -I . ^
  --go_out ../service --micro_out paths=source_relative ^
  --go-grpc_out ../service --go-grpc_opt paths=source_relative ^
  --grpc-gateway_out ../service --grpc-gateway_opt paths=source_relative ^
  ./comment_list.proto