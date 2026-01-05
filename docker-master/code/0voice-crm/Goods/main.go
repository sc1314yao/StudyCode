package main

import (
	"Goods/config"
	"Goods/proto"
	"Goods/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	conf := config.GetConf()
	proto.RegisterGoodsServer(s, server.NewGoodsServer(conf))

	healthCheck := health.NewServer()
	grpc_health_v1.RegisterHealthServer(s, healthCheck)
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
