package server

import (
	"context"
	"etcd/echo"
	"etcd/pkg/etcd"
	"etcd/pkg/lock"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"math/rand"
	"time"
)

type echoServer struct {
	echo.UnimplementedEchoServer
}

func NewEchoServer() echo.EchoServer {
	return &echoServer{}
}
func (s *echoServer) UnaryEcho(ctx context.Context, in *echo.EchoRequest) (*echo.EchoResponse, error) {
	fmt.Println("server recv: ", in.Message)
	cli, err := etcd.GetClient()
	if err != nil {
		return nil, err
	}
	//创建租约
	leaseRes, err := cli.Grant(context.Background(), 30)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//添加数据
	cli.Put(context.Background(), "/0voice/key1", "value1", clientv3.WithLease(leaseRes.ID))
	cli.Put(context.Background(), "/0voice/key2", "value2", clientv3.WithLease(leaseRes.ID))
	cli.Put(context.Background(), "/0voice/key3", "value3", clientv3.WithLease(leaseRes.ID))

	//查询
	getRes, err := cli.Get(context.Background(), "/0voice/key", clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, kv := range getRes.Kvs {
		fmt.Printf("key: %s,value: %s,leaseID:%x \n", kv.Key, kv.Value, kv.Lease)
	}
	//删除
	cli.Delete(context.Background(), "/0voice/key", clientv3.WithPrefix())
	//事务
	txn := cli.Txn(context.Background())
	// 支持的条件
	// clientv3.ModRevision("/0voice/key4") 获取对应key的最近一次修改的版本号
	// clientv3.Version("/0voice/key4") 获取key被修改的次数
	// clientv3.Value("/0voice/key4") 获取key的value值
	// clientv3.LeaseValue("/0voice/key4") 获取key对应的租约ID，若没有租约则返回0
	txn.If(
		clientv3.Compare(clientv3.CreateRevision("/0voice/key4"), "=", 0),
	).Then(
		clientv3.OpPut("/0voice/key4", "value4"),
		clientv3.OpPut("/0voice/key5", "value5"),
	).Else(
		clientv3.OpDelete("/0voice/key4"),
		clientv3.OpDelete("/0voice/key5"),
	).Commit()

	return &echo.EchoResponse{
		Message: "hello client",
	}, nil
}

func (s *echoServer) UnaryLock(ctx context.Context, in *echo.EchoRequest) (*echo.EchoLockResponse, error) {
	fmt.Println("server recv: ", in.Message)
	mutex, err := lock.NewMutex("mylock", 10)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = mutex.Lock(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	randID := rand.Int63()
	start := time.Now()
	fmt.Printf("ID: %d 上锁成功，time：%v \n", randID, start.Unix())
	time.Sleep(time.Second * 5)
	end := time.Now()
	err = mutex.Unlock(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fmt.Printf("ID: %d 解锁成功，time：%v \n", randID, start.Unix())

	return &echo.EchoLockResponse{
		Id:      randID,
		Message: "hello client",
		Start:   timestamppb.New(start),
		End:     timestamppb.New(end),
	}, nil
}
