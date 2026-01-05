package register

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type cusRegister struct {
	client *clientv3.Client
}

func NewCusRegister(client *clientv3.Client) ServiceRegister {
	return &cusRegister{
		client: client,
	}
}
func (r *cusRegister) Register(scheme, serviceName, serviceAddr string) error {
	ctx := context.Background()
	leaseRes, err := r.client.Grant(ctx, 10)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("/%s/%s/%s", scheme, serviceName, serviceAddr)
	_, err = r.client.Put(ctx, key, serviceAddr, clientv3.WithLease(leaseRes.ID))
	if err != nil {
		return err
	}
	chKeepaliveResp, err := r.client.KeepAlive(ctx, leaseRes.ID)
	if err != nil {
		return err
	}
	go func() {
		for resp := range chKeepaliveResp {
			fmt.Printf("ID: %x,ttl:%d \n", resp.ID, resp.TTL)
		}
	}()

	return nil
}
