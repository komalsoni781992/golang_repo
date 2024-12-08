package main

import (
	"fmt"
	"strconv"
	"sync"
)

/*q3. Create a bookCab function that takes the name of the user trying to book a cab
  Assume only one user can book a cab at a time
  Create a global variable to hold the number of cabs available

  Check if a cab is available, if yes print a msg cab is available otherwise unavailable

  In the main function, run a loop 5 times and call bookCab function as goroutine to simulate
  multiple users are trying to book a cab*/

// Create a global variable to hold the number of cabs available
var numCabs int64 = 2

func main() {

	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go bookCab("user: "+strconv.Itoa(i), m, wg)
	}
	wg.Wait()
}

// bookCab - Create a bookCab function that takes the name of the user trying to book a cab
func bookCab(user string, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	defer m.Unlock()
	m.Lock()
	if numCabs > 0 {
		numCabs--
		fmt.Println("Cab is available. user:", user, " got lucky!!")
	} else {
		fmt.Println("Better luck next time user:", user)
	}
}
