@REM goctl model mysql ddl -src ./model/newProduct.sql -dir ./model -c
@REM goctl api go -api ./api/newProduct.api -dir ./api
@REM goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

go run product.go -f etc/product.yaml
