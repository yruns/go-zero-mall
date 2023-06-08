@REM goctl model mysql ddl -src ./model/product.sql -dir ./model -c
@REM goctl api go -api ./api/product.api -dir ./api
@REM goctl rpc protoc product.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.
