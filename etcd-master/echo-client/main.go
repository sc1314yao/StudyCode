package main

import (
	"etcd/echo-client/client"
	"etcd/pkg/discovery"
	"etcd/pkg/etcd"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var (
	scheme      = flag.String("scheme", "myscheme", "")
	serviceName = flag.String("service_name", "", "")
)

func main() {
	flag.Parse()
	addr := ""
	if *scheme != "" && *serviceName != "" {
		addr = fmt.Sprintf("%s:///%s", *scheme, *serviceName)
		cli, _ := etcd.GetClient()
		ds := discovery.NewCusDiscovery(cli, *scheme, *serviceName)
		fmt.Println(ds.Discovery())
	}
	conn, err := grpc.Dial(addr, grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`), grpc.WithResolvers(&client.MySchemeBuilder{}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 5; i++ {
		client.CallEcho(conn)
		client.CallEchoLock(conn)
	}
}
