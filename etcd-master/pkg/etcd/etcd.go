package etcd

import clientv3 "go.etcd.io/etcd/client/v3"

const PROXY = "proxy"
const GATEWAY = "gateway"

func GetClient() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: getEtcdEndpoints(GATEWAY),
	})
	return cli, err
}

func getEtcdEndpoints(typ string) []string {
	endpoints := []string{"192.168.239.161:2379", "192.168.239.161:12379", "192.168.239.161:22379"}
	if typ == PROXY {
		endpoints = []string{"192.168.239.161:23791", "192.168.239.161:23792", "192.168.239.161:23793"}
	} else if typ == GATEWAY {
		endpoints = []string{"192.168.239.161:23790"}
	}
	return endpoints
}
