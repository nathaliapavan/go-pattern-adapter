package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
)

type MemcachedAdapter struct {
	client *memcache.Client
}

func NewMemcachedAdapter() *MemcachedAdapter {
	mc := memcache.New("memcached:11211")
	return &MemcachedAdapter{client: mc}
}

func (m *MemcachedAdapter) Set(key string, value string) error {
	return m.client.Set(&memcache.Item{Key: key, Value: []byte(value)})
}

func (m *MemcachedAdapter) Get(key string) (string, error) {
	item, err := m.client.Get(key)
	if err != nil {
		return "", err
	}
	return string(item.Value), nil
}

func (m *MemcachedAdapter) Delete(key string) error {
	return m.client.Delete(key)
}
