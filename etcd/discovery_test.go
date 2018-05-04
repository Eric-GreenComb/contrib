package etcd

import (
	"fmt"
	"testing"
	// "log"
	// "time"
	// "github.com/coreos/etcd/client"
)

func TestBasic(t *testing.T) {

	// cfg := client.Config{
	// 	Endpoints:               []string{"http://127.0.0.1:2379"},
	// 	Transport:               client.DefaultTransport,
	// 	HeaderTimeoutPerRequest: time.Second,
	// }
	// c, err := client.New(cfg)

	// if err != nil {
	// 	log.Panicln(err)
	// }

	// kapi := client.NewKeysAPI(c)

	// client := EtcdReigistryClient{
	// 	EtcdRegistryConfig{
	// 		ServiceName:  "banerwai/micros/query/category/addr",
	// 		InstanceName: "127.0.0.1:8080",
	// 		BaseURL:      "127.0.0.1:8080",
	// 	},
	// 	kapi,
	// }

	client := ReigistryClient{
		RegistryConfig{
			ServiceName:  "/banerwai/service/test",
			InstanceName: "127.0.0.1:8080",
			BaseURL:      "127.0.0.1:8080",
		},
		KeysAPI,
	}
	client.Register()

	response, _ := client.ServicesByName("banerwai/service/test")
	if len(response) == 0 {
		t.Error("No service registered")
	}
	fmt.Println(response)
	// client.Unregister()
	// response, _ = client.ServicesByName("test")
	// fmt.Println(len(response))
	// if len(response) != 0 {
	// 	t.Error("Service not  unregistered")
	// }
}

// func TestKeepAlive(t *testing.T) {

// 	cfg := client.Config{
// 		Endpoints:               []string{"http://127.0.0.1:2379"},
// 		Transport:               client.DefaultTransport,
// 		HeaderTimeoutPerRequest: time.Second,
// 	}
// 	c, err := client.New(cfg)

// 	if err != nil {
// 		log.Panicln(err)
// 	}

// 	kapi := client.NewKeysAPI(c)

// 	client := EtcdReigistryClient{
// 		EtcdRegistryConfig{
// 			ServiceName:  "test",
// 			InstanceName: "test1",
// 			BaseURL:      "127.0.0.1:8080",
// 		},
// 		kapi,
// 	}

// 	client.Register()

// 	time.Sleep(50 * time.Second)

// 	response, _ := client.ServicesByName("test")
// 	fmt.Println(response)
// 	log.Println(response)

// 	// Ahoj
// }
