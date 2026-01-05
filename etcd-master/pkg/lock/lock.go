package lock

import (
	"etcd/pkg/etcd"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func NewMutex(prefix string, ttl int) (*concurrency.Mutex, error) {
	cli, err := etcd.GetClient()
	if err != nil {
		return nil, err
	}
	session, err := concurrency.NewSession(cli, concurrency.WithTTL(ttl))
	if err != nil {
		return nil, err
	}
	mutex := concurrency.NewMutex(session, prefix)
	return mutex, nil
}
