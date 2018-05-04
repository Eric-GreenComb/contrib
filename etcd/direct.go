package etcd

import (
	"github.com/coreos/etcd/client"
	"golang.org/x/net/context"
)

// Set etcd direct set a key/value
func Set(key, value string) (*client.Response, error) {
	return KeysAPI.Set(context.Background(), key, value, nil)
}

// Get etcd derict get a etcd response by key
func Get(key string) (*client.Response, error) {
	return KeysAPI.Get(context.Background(), key, nil)
}

// GetValue etcd derict get a value by key
func GetValue(key string) (string, error) {
	resp, err := KeysAPI.Get(context.Background(), key, nil)
	if err != nil {
		return "", err
	}
	return resp.Node.Value, nil
}

// GetString etcd derict get a string by key
func GetString(key string) string {
	resp, err := GetValue(key)
	if err != nil {
		return ""
	}
	return resp
}
