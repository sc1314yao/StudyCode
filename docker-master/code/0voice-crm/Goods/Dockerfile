FROM golang:1.20 as stage0
RUN go env -w GOPROXY=https://proxy.golang.com.cn,https://goproxy.cn,direct
ADD ./ /src/goods
WORKDIR /src/goods
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o goods .

FROM alpine:3.18 as stage1
ADD ./grpc_health_probe-linux-amd64 /usr/bin/grpc_health_probe
RUN chmod +x /usr/bin/grpc_health_probe
MAINTAINER nick
WORKDIR /app/
ADD ./config.yaml /app/config.yaml
COPY --from=stage0 /src/goods/goods ./
ENTRYPOINT ["./goods"]