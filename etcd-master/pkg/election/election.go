package election

import (
	"etcd/pkg/etcd"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func NewElection(prefix string, ttl int) (*concurrency.Election, error) {
	cli, err := etcd.GetClient()
	if err != nil {
		return nil, err
	}
	session, err := concurrency.NewSession(cli, concurrency.WithTTL(ttl))
	if err != nil {
		return nil, err
	}
	elect := concurrency.NewElection(session, prefix)
	return elect, nil
}
