@REM goctl model mysql ddl -src ./model/newProduct.sql -dir ./model -c
@REM goctl api go -api ./api/newProduct.api -dir ./api
@REM goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

go run product.go -f etc/product.yaml

@REM 启动product rpc
go run product/rpc/product.go -f product/rpc/etc/product.yaml
@REM 启动product api
go run product/api/product.go -f product/api/etc/product.yaml

@REM 启动user rpc
go run user/rpc/user.go -f user/rpc/etc/user.yaml
@REM 启动user api
go run user/api/user.go -f user/api/etc/user.yaml

@REM 启动order rpc
go run order/rpc/order.go -f order/rpc/etc/order.yaml
@REM 启动order api
go run order/api/order.go -f order/api/etc/order.yaml

