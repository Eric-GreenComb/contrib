package cache

import (
	"sync"
)

type memoryCache struct {
	lock  *sync.RWMutex
	items map[interface{}]interface{}
}

func (mc *memoryCache) set(key interface{}, value interface{}) error {
	mc.lock.Lock()
	defer mc.lock.Unlock()
	mc.items[key] = value
	return nil
}

func (mc *memoryCache) get(key interface{}) interface{} {
	mc.lock.RLock()
	defer mc.lock.RUnlock()

	if val, ok := mc.items[key]; ok {
		return val
	}
	return nil
}

// GCache GCache
var GCache = &memoryCache{
	lock:  new(sync.RWMutex),
	items: make(map[interface{}]interface{}),
}
