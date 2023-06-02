package main

import (
	"fmt"
	"udemyHT1w0/cache"
)

func main() {
	cacheExample := cache.NewCache()

	err := cacheExample.Set("testkey1", "testvalue1")
	err = cacheExample.Set("key2", 0)
	err = cacheExample.Set("testkey3", 42)
	err = cacheExample.Set("testkey5", " ")
	err = cacheExample.Set("key6", "")
	err = cacheExample.Set("key4", nil)
	if err == nil {
		fmt.Println("cache consists:", cacheExample)
	} else {
		fmt.Println("err:", err)
	}

	temp, err := cacheExample.Get("testkey3")
	fmt.Println(temp)

	temp, err = cacheExample.Get("testkey4")
	fmt.Println(err)
}
