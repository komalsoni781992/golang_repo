package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := range 10 {
		wg.Add(1)
		go work(int64(i+20), wg)
	}
	wg.Wait()
}

func work(id int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Printf("work {%d} is going on.\n", id)
	time.Sleep(1 * time.Second)
}
