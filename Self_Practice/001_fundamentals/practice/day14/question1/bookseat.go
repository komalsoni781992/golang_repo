package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

/*Make a struct Theater with the following fields: Seats(int=1), RWMutex, userName chan string.

   Create two methods over a struct

   The first method book a seat in the theater. If the seat is equal to zero, no one can book it.
   ( In the booking method, put simple print statements that show booking has been made if seats are available)

   Once the seat is booked in Theater, add the name of the user in the userName channel.
   Create a second Method, printInvoice(),  It fetches the userName from the channel and print it.

  Note:-
   Call the bookSeats & printInvoice() method as a goroutine in the main function.
   For example:-

   for i:=1; i<=3; i++ {
        go t.bookSeats()
   }
   go t.printInvoice()

   The program should quit gracefully without deadlock.*/

type Seats int

type Theater struct {
	seats    int
	m        *sync.RWMutex
	username chan string
}

func (t *Theater) bookSeat(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("welcome to the website", name)
	func() {
		t.m.Lock()
		defer t.m.Unlock()
		fmt.Println("Booking seat for: ", name)
		if t.seats >= 1 {
			fmt.Println("seat is available for", name)
			time.Sleep(5 * time.Second)
			fmt.Println("booking confirmed", name)
			t.seats--
			t.username <- name
		} else {
			fmt.Println("car is not available for: ", name)
		}
	}()
}

func (t *Theater) printInvoice(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Printing invoice for: ")
	for v := range t.username {
		fmt.Println(v)
	}
}

func main() {
	rwm := new(sync.RWMutex)
	t := Theater{seats: 1, m: rwm, username: make(chan string)}
	wg := new(sync.WaitGroup)
	workerWg := new(sync.WaitGroup)
	for i := 1; i <= 3; i++ {
		workerWg.Add(1)
		go t.bookSeat("user: "+strconv.Itoa(i), workerWg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		workerWg.Wait()
		close(t.username)
	}()
	go t.printInvoice(wg)
	wg.Wait()
}
