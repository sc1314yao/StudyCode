package client

import (
	"etcd/pkg/discovery"
	"etcd/pkg/etcd"
	"google.golang.org/grpc/resolver"
	"log"
)

const SCHEME = "myscheme"
const SERVICENAME = "echo"

type MySchemeBuilder struct {
}

func (b *MySchemeBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &MyResolver{
		target: target,
		cc:     cc,
	}
	r.start()
	return r, nil
}
func (b *MySchemeBuilder) Scheme() string {
	return SCHEME
}

type MyResolver struct {
	target resolver.Target
	cc     resolver.ClientConn
}

func (r *MyResolver) start() {
	client, err := etcd.GetClient()
	if err != nil {
		log.Println(err)
		return
	}
	d := discovery.NewCusDiscovery(client, SCHEME, SERVICENAME)
	addrList := d.Discovery()
	if len(addrList) == 0 {
		log.Println("没有发现任何地址")
		return
	}
	addrs := make([]resolver.Address, len(addrList))
	for i, s := range addrList {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (r *MyResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (r *MyResolver) Close()                                {}
