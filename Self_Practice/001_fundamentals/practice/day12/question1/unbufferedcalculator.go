package main

import (
	"fmt"
	"sync"
)

/*
q4. Create 4 functions

	Add(int,int),Sub(int,int),Divide(int,int), CollectResults()
	Add,Sub,Divide do their operations and send value to an unbuffered channel

	CollectResult() -> It would receive the values from the channel and print it
*/
//https://github.com/Diwakarsingh052/genspark/blob/main/12-concurrency/8-unbuf-chan.go

func main() {
	var c chan int = make(chan int)
	wg := new(sync.WaitGroup)
	for i := 0; i <= 3; i++ {
		wg.Add(4)
		go Add(10, i, &c, wg)
		go Sub(10, i, &c, wg)
		go Divide(10, i, &c, wg)
		go CollectResults(&c, wg)
		wg.Wait()
		fmt.Println("------------------------------")
	}
}

func Add(a int, b int, c *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Adding")
	*c <- a + b
}

func Sub(a int, b int, c *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Subtracting")
	*c <- a - b
}

func Divide(a int, b int, c *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Dividing")
	if b == 0 {
		fmt.Println("Denominator was sent to be zero")
		*c <- 0
		return
	}
	*c <- a / b
}

func CollectResults(c *chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Collecting")
	for i := 1; i <= 3; i++ {
		fmt.Println(<-*c)
	}
}
