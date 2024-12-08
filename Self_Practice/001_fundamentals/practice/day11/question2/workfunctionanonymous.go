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
		go yetAnotherWork(int64(i+20), wg)
	}
	wg.Wait()
}

/*func work(id int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Printf("work {%d} is going on.\n", id)
	time.Sleep(1 * time.Second)

	awg := new(sync.WaitGroup)
	awg.Add(1)
	go func() {
		defer awg.Done()
		fmt.Printf("I think I am little anonymous here. :) : {%d}\n", id)
		time.Sleep(100 * time.Millisecond)
	}()
	awg.Wait()
}*/

func yetAnotherWork(id int64, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Printf("work {%d} is going on.\n", id)
	time.Sleep(1 * time.Second)
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("I think I am little anonymous here. :) : {%d}\n", id)
		time.Sleep(100 * time.Millisecond)
	}()
}
