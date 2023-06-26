package cache

import (
	"errors"
	"sync"
	"time"
)

type value struct {
	value interface{}
	ttl   *time.Time
}

type Cache struct {
	ticker *time.Ticker
	data   sync.Map
	ttl    time.Duration
}

func NewCache(ttl time.Duration) *Cache {
	newcache := &Cache{
		ticker: time.NewTicker(time.Second * 1),
		data:   sync.Map{},
		ttl:    ttl,
	}
	go newcache.backgroundCacheCleaner()

	return newcache
}

// background goroutine to clean up expired keys in the cache
func (obj *Cache) backgroundCacheCleaner() {
	for {
		<-obj.ticker.C
		obj.data.Range(func(key, v interface{}) bool {
			vv, ok := v.(*value)
			if !ok {
				return true
			}

			if vv.ttl == nil {
				return true
			}

			if time.Now().After(*vv.ttl) {
				obj.data.Delete(key)
			}

			return true
		})
	}
}

func (obj *Cache) Set(key string, v interface{}) error {
	t := time.Now().Add(obj.ttl)
	obj.data.Store(key, &value{v, &t})
	return nil
}
func (obj *Cache) Get(key string) (interface{}, error) {
	load, resultFlag := obj.data.Load(key)

	if !resultFlag {
		return nil, errors.New("no such key exists")
	}
	vv, ok := load.(*value)
	if !ok {
		return nil, errors.New("no such key exists")
	}

	return vv.value, nil
}
func (obj *Cache) Delete(key string) error {
	obj.data.Delete(key)
	return nil
}
