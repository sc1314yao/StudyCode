# 商品服务
## grpc 代码生成
``` 
protoc.exe --go_out=. --go_opt=paths=source_relative .\proto\goods.proto 
protoc.exe --go-grpc_out=. --go-grpc_opt=paths=source_relative .\proto\goods.proto
```
## 依赖安装
``` 
go mod tidy
```
## 其他依赖
``` 
go get github.com/spf13/viper@v1.16.0
```
## docker 镜像构建
``` 
docker build -t goods:v1.0.0 .
```