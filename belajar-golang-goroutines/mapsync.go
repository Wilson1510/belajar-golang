package main

import (
	"fmt"
	"sync"
	"reflect"
)

func addToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	defer group.Done()
	group.Add(1)
	data.Store(value, value)
}

func main() {
	var data sync.Map
	var group sync.WaitGroup

	for i := 0; i < 100; i++ {
		go addToMap(&data, i, &group)
	}

	group.Wait()
	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, value, reflect.TypeOf(key), reflect.TypeOf(value))
		return true
	})
}