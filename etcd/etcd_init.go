package etcd

import (
	"flag"
	"log"
	"time"

	"github.com/coreos/etcd/client"
)

// KeysAPI etcd KeysAPI
var KeysAPI client.KeysAPI

const (
	// TTL is a time to live
	// for record in etcd
	TTL = 30 * time.Second

	// KeepAlivePeriod is period of
	// goroutine to
	// refresh the record in etcd.
	KeepAlivePeriod = 20 * time.Second
)

func init() {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		etcdAddr = fs.String("etcd.addr", "http://127.0.0.1:2379", "Address for etcd server")
	)

	cfg := client.Config{
		Endpoints:               []string{*etcdAddr},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		log.Panicln(err)
	}
	KeysAPI = client.NewKeysAPI(c)
}
