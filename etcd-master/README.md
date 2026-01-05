# grpc 代码生成
``` 
protoc --go_out=. --go_opt=paths=source_relative echo/echo.proto
protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative echo/echo.proto
```

# 依赖对齐
``` 
go mod tidy
```
# 安装etcdv3的go客户端
``` 
go get go.etcd.io/etcd/client/v3
```