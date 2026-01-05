FROM golang:1.20 as stage0
RUN go env -w GOPROXY=https://proxy.golang.com.cn,https://goproxy.cn,direct
ADD ./ /src/0voiceGateway
WORKDIR /src/0voiceGateway
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o 0voiceGateway .

FROM alpine:3.18 as stage1
ADD ./curl-amd64 /usr/bin/curl
RUN chmod +x /usr/bin/curl
MAINTAINER nick
WORKDIR /app/
ADD ./config.yaml /app/config.yaml
COPY --from=stage0 /src/0voiceGateway/0voiceGateway ./
ENTRYPOINT ["./0voiceGateway"]