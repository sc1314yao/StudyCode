# 0voiceGateway
网关对外提供统一接口管理
## 依赖安装
``` 
go get -u github.com/gin-gonic/gin@v1.9.1
go get github.com/spf13/viper@v1.16.0
```
## 生成grpc存根
``` 
protoc.exe --go_out=. --go_opt=paths=source_relative .\services\goods\proto\goods.proto 
protoc.exe --go-grpc_out=. --go-grpc_opt=paths=source_relative .\services\goods\proto\goods.proto
```

## 镜像构建
``` 
docker build -t 0voice-gateway:v1.0.0 .
```