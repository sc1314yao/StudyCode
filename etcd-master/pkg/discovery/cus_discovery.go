package discovery

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"sync"
)

type serviceAddrCache struct {
	data map[string]map[string]string
	lock sync.RWMutex
}
type sdInstanceCache struct {
	data map[string]*cusDiscovery
	lock sync.RWMutex
}

var addrCache *serviceAddrCache
var sdCache *sdInstanceCache

func init() {
	addrCache = &serviceAddrCache{
		data: map[string]map[string]string{},
		lock: sync.RWMutex{},
	}
	sdCache = &sdInstanceCache{
		data: map[string]*cusDiscovery{},
		lock: sync.RWMutex{},
	}
}

type cusDiscovery struct {
	client      *clientv3.Client
	scheme      string
	serviceName string
}

func NewCusDiscovery(client *clientv3.Client, scheme, serviceName string) ServiceDiscovery {
	key := getPrefix(scheme, serviceName)
	sdCache.lock.Lock()
	defer sdCache.lock.Unlock()
	d, ok := sdCache.data[key]
	if !ok {
		d = &cusDiscovery{
			client:      client,
			scheme:      scheme,
			serviceName: serviceName,
		}
		sdCache.data[key] = d
		d.loadAndWatch()
	}
	return d
}

func (d *cusDiscovery) Discovery() []string {
	list := make([]string, 0)
	key := getPrefix(d.scheme, d.serviceName)
	addrCache.lock.RLock()
	defer addrCache.lock.RUnlock()
	mp, ok := addrCache.data[key]
	if !ok {
		return list
	}
	for _, v := range mp {
		list = append(list, v)
	}
	return list
}

func (d *cusDiscovery) loadAndWatch() {
	prefix := getPrefix(d.scheme, d.serviceName)
	//第一次 get 数据
	getRes, err := d.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return
	}
	mp := make(map[string]string, getRes.Count)
	for _, kv := range getRes.Kvs {
		k := string(kv.Key)
		v := string(kv.Value)
		mp[k] = v
	}
	addrCache.lock.Lock()
	defer addrCache.lock.Unlock()
	addrCache.data[prefix] = mp

	// go watch
	go d.watch()
}

func (d *cusDiscovery) watch() {
	// 根据事件，修改本地缓存
	prefix := getPrefix(d.scheme, d.serviceName)
	rch := d.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wres := range rch {
		for _, event := range wres.Events {
			addrCache.lock.Lock()
			if event.Type == clientv3.EventTypeDelete {
				k := string(event.Kv.Key)
				delete(addrCache.data[prefix], k)
			}
			if event.Type == clientv3.EventTypePut {
				k := string(event.Kv.Key)
				v := string(event.Kv.Value)
				addrCache.data[prefix][k] = v
			}
			addrCache.lock.Unlock()
		}
	}
}

func getPrefix(scheme, serviceName string) string {
	return fmt.Sprintf("/%s/%s/", scheme, serviceName)
}
