package main

import (
	"context"
	"etcd/echo"
	"etcd/echo-server/server"
	election2 "etcd/pkg/election"
	"etcd/pkg/etcd"
	register2 "etcd/pkg/register"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

var (
	scheme      = flag.String("scheme", "myscheme", "")
	serviceName = flag.String("service_name", "", "")
	nodeName    = flag.String("node_name", "", "")
	port        = flag.Int("port", 50051, "")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	echoServer := server.NewEchoServer()
	echo.RegisterEchoServer(s, echoServer)

	if *serviceName != "" && *nodeName != "" {
		go electLeader(*serviceName, *nodeName)
	}
	if *scheme != "" && *serviceName != "" {
		serviceRegister(*scheme, *serviceName, fmt.Sprintf("%s:%d", "localhost", *port))
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}

func electLeader(serviceName, nodeName string) error {
	electName := "/elect/" + serviceName
	election, err := election2.NewElection(electName, 10)
	if err != nil {
		return err
	}
	go func() {
		ch := election.Observe(context.Background())
		for res := range ch {
			fmt.Printf("监听leader选举情况：key: %s, value: %s \n", res.Kvs[0].Key, res.Kvs[0].Value)
		}
	}()
	err = election.Campaign(context.Background(), nodeName)
	if err != nil {
		return err
	}
	// TODO 业务操作，成为leader的节点，应该做的一些初始化动作
	fmt.Printf("%v 当选为leader节点 \n", nodeName)
	time.Sleep(time.Second * 30)
	// 卸任
	err = election.Resign(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("%v 卸任leader \n", nodeName)
	return nil
}

func serviceRegister(scheme, serviceName, serviceAddr string) error {
	cli, err := etcd.GetClient()
	if err != nil {
		return err
	}
	register := register2.NewCusRegister(cli)
	err = register.Register(scheme, serviceName, serviceAddr)
	return err
}
