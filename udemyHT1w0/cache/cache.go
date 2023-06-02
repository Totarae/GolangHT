package cache

import "errors"

type Cache struct {
	valuesMap map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{make(map[string]interface{})}
}

func (obj *Cache) Set(key string, value interface{}) {

}
func (obj *Cache) Get(key string) (interface{}, error) {
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
