package concurrency

//https://github.com/Diwakarsingh052/genspark/blob/main/12-concurrency/2-go.go
import (
	"fmt"
	"sync"
	"time"
)

// TestConcurrencyWithWaitGroup - tests wait group functionality
func TestConcurrencyWithWaitGroup() {
	// using waitgroup keep track of goroutines spun up
	wg := new(sync.WaitGroup) // waitgroup must be a pointer

	// waitgroup counter represents number of goroutine we are running

	wg.Add(1) // add 1 to the counter
	go func() {
		defer wg.Done() // decrement the counter
		time.Sleep(3 * time.Second)
		fmt.Println("Hello, world!")
	}()

	fmt.Println("some work going on in the main function")
	wg.Wait() // wait until the counter is not 0

	fmt.Println("Done!")

}
