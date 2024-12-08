package main

// create a data-race for the map and detect with race detector

import (
	"fmt"
	"strconv"
	"sync"
)

var user map[int]string = make(map[int]string, 5)

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go updateUser(1, "H: "+strconv.Itoa(i), wg, m)
	}
	wg.Wait()
}

func updateUser(key int, val string, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	//m.Lock()
	//defer m.Unlock()
	user[key] = val
	fmt.Println(user[key])
}
