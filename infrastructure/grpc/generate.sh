 protoc -I ./proto  --go_out ./proto --go_opt paths=source_relative \
 --go-grpc_out ./proto --go-grpc_opt paths=source_relative    ./proto/ports/ports.proto

     protoc -I ./proto \
       --go_out ./proto --go_opt paths=source_relative \
       --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
       --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
       --openapiv2_out ./openapiv2 --openapiv2_opt logtostderr=true  \
       ./proto/ports/ports.proto

