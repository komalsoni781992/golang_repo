package concurrency

//https://github.com/Diwakarsingh052/genspark/blob/main/12-concurrency/9-buffered-chan.go

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/spec#Send_statements
// A send on a buffered channel can proceed if there is room in the buffer.

// TestBufferedChannel - introduction to bufferd channels
func TestBufferedChannel() {
	wg := new(sync.WaitGroup)
	// try to play with this number, to see how results changes
	// try to make it as an unbuffered chan as well
	/*

		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		wg.Add(1)
		go func()  {
			defer wg.Done()
			ch  <- 10
		}()

		wg.Add(1)
		go func()  {
			defer wg.Done()
			fmt.Println(<-ch)
		}()
		wg.Wait()

	*/

	/*
		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		wg.Add(1)
		go func()  {
			defer wg.Done()
			ch  <- 10
			ch  <- 10
		}()

		wg.Add(1)
		go func()  {
			defer wg.Done()
			fmt.Println(<-ch)
		}()
		wg.Wait()
	*/

	/*
		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		for i:= 1; i <= 2; i++ {
			wg.Add(1)
			go func()  {
				defer wg.Done()
				ch  <- 10
				ch  <- 10
			}()
		}

		wg.Add(1)
		go func()  {
			defer wg.Done()
			fmt.Println(<-ch)
		}()
		wg.Wait()
	*/

	/*
		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		for i:= 1; i <= 3; i++ {
			wg.Add(1)
			go func()  {
				defer wg.Done()
				ch  <- 10
				ch  <- 10
			}()
		}

		wg.Add(1)
		go func()  {
			defer wg.Done()
			fmt.Println(<-ch)
		}()
		wg.Wait()
	*/

	/*
		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		for i:= 1; i <= 3; i++ {
			wg.Add(1)
			go func()  {
				defer wg.Done()
				ch  <- 10
				ch  <- 10
			}()
		}

		for i:= 1; i <= 2; i++ {
			fmt.Println(<-ch)
		}
		wg.Wait()
	*/

	/*
		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		for i:= 1; i <= 2; i++ {
			fmt.Println(<-ch)
		}

		for i:= 1; i <= 3; i++ {
			wg.Add(1)
			go func()  {
				defer wg.Done()
				ch  <- 10
				ch  <- 10
			}()
		}

		wg.Wait()
	*/

	/*
		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		wg.Add(1)
		go func()  {
			for i:= 1; i <= 2; i++ {
				fmt.Println(<-ch)
			}
		}

		for i:= 1; i <= 3; i++ {
			wg.Add(1)
			go func()  {
				defer wg.Done()
				ch  <- 10
				ch  <- 10
			}()
		}

		wg.Wait()
	*/

	/*
		wg := new(sync.WaitGroup)
		ch := make(chan int, 2)
		wg.Add(1)
		go func()  {
			defer wg.Done()
			for i:= 1; i <= 2; i++ {
				fmt.Println(<-ch)
			}
		}

		for i:= 1; i <= 3; i++ {
			wg.Add(1)
			go func()  {
				defer wg.Done()
				ch  <- 10
				ch  <- 10
			}()
		}

		wg.Wait()
	*/

	ch := make(chan int, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			// when we recv value we make one slot empty in the buffer, and more value could be sent over it
			// make sure to recv all the values from the sender, no guarantees given by buffered chan
			time.Sleep(2 * time.Second)
			fmt.Println(<-ch)
		}
	}()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- i // it would only see if there is room in the buffer, if yes it would send the value
			// buf chan doesn't care about recv
			fmt.Println("sent", i)
		}()
	}

	wg.Wait()
	fmt.Println("end of main")

}
