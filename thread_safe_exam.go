package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	Data map[string]interface{}
	*sync.Mutex
}

func (m *SafeMap) put(key string, value interface{},wg *sync.WaitGroup) {
	//defer wg.Done()
	m.Lock()
	defer m.Unlock()
	m.Data[key] = value
}

func (m *SafeMap) get(key string, wg *sync.WaitGroup) (interface{}, bool) {
	//defer wg.Done()
	m.Lock()
	m.Unlock()
	v, ok := m.Data[key]
	return v, ok
}

//func newSafeMap(data map[string]interface{}) *SafeMap{
//	return &SafeMap{data,&sync.Mutex{}}
//}

func main() {
	d := &SafeMap{
		map[string]interface{}{
			"key1": "v1",
			"key2": "v2",
		},
		&sync.Mutex{},
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		d.get("key1",wg)
		defer wg.Done()
	}()
	go func() {
		d.put("key1","key1 update",wg)
		defer wg.Done()
	}()
	wg.Wait()
	fmt.Println(d.Data)
}
