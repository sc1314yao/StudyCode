# 服务部署
### 创建网络
``` 
docker network create -d overlay \
--attachable \
--subnet=10.12.0.0/16 \
--gateway=10.12.0.100 \
0voice-overlay
```
## 部署goods
```
cd Goods
```
### 构建镜像
``` 
docker build -t 192.168.239.142:5000/0voice-goods:v1.0.0 .
docker push 192.168.239.142:5000/0voice-goods:v1.0.0
```
### 创建配置文件
``` 
docker config create goods-conf config.yaml
```
### 部署服务
``` 
docker service create --name 0voice-goods -p 50052:50051 \
--network 0voice-overlay \
--config src=goods-conf,target=/app/config.yaml \
--replicas 3 --limit-cpu 0.3 --reserve-cpu 0.1 \
--update-parallelism=2 \
--health-cmd "grpc_health_probe -addr=:50051" \
--health-interval 5s --health-retries 3 \
--with-registry-auth \
192.168.239.142:5000/0voice-goods:v1.0.0
```

## 部署gateway
### 构建镜像
``` 
cd 0voiceGateway
docker build -t 192.168.239.142:5000/0voice-gateway:v1.0.0 .
docker push 192.168.239.142:5000/0voice-gateway:v1.0.0
```
### 创建配置文件
``` 
docker config create gateway-conf ./0voiceGateway/config.yaml
```
### 部署服务
``` 
docker service create --name 0voice-gateway -p 8082:8081 \
--network 0voice-overlay \
--config src=gateway-conf,target=/app/config.yaml \
--replicas 3 --limit-cpu 0.3 --reserve-cpu 0.1 \
--update-parallelism=2 \
--health-cmd "curl -f http://localhost:8081/api/health" \
--health-interval 5s --health-retries 3 \
--with-registry-auth \
192.168.239.142:5000/0voice-gateway:v1.0.0
```

### stack部署和更新
``` 
docker stack deploy -c compose-stack.yaml --with-registry-auth 0voicestack
```