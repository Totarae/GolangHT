package cache

import "errors"

type Cache struct {
	valuesMap map[string]any
}

func NewCache(valuesMap map[string]any) *Cache {
	return &Cache{make(map[string]any)}
}

func (obj *Cache) Set(key string, value interface{}) {

}
func (obj *Cache) Get(key string) (any, error) {
	value, exists := obj.valuesMap[key]

	if exists {
		return value, nil
	}
	return nil, errors.New("no such key exists")
}
func (obj *Cache) Delete(key string) error {
	delete(obj.valuesMap, key)
	return nil
}
