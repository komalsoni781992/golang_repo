package concurrency

// https://github.com/Diwakarsingh052/genspark/blob/main/12-concurrency/1-go.go

// parallelism(doing multiple things at once - happens by default for multiple cores)
// versus concurrency (dealing with multiple things at once - efficiency)

// Multiple  Go routines(KB) are multiplexed over kernel level thread(MB)
// By default, go routines are configured in Go
// First time kernel level thread is created and multiple go routines remain in its local run queue (LRQ)
// First goroutine created is Main goroutine
// If a kernel thread is overcrowded then new thread is created
// Go scheduler also takes care to steal work from one thread and assign to idle thread in case of parallelism
// If thread is idle, it is not killed, it is kept alive in sleep state to be used for later time

// If a goroutines makes a system call and goes in blocked state, the kernel level thread also goes in blocked state
// For the remaining goroutines in ready queue of the blocked thread, a new kernel level thread is created and assigned to it
// The blocked thread when becomes idle is kept alive to be used later

// functions and method calls can be made concurrent

// There is a global run queue(GRQ) shared among all kernel level Go threads, anything in this queue has low priority
// than those in LRQ(long running methods)

import (
	"fmt"
	"time"
)

// concurrency is dealing with a lot of things at once
// parallelism is doing multiple things at once

// Concurrency is not Parallelism by Rob Pike
// https://www.youtube.com/watch?v=oV9rvDllKEg&ab_channel=gnbitcom

// TestConcurrency - TestConcurrency - tests concurrency
func TestConcurrency() {
	//panic("panic") // reveals goroutine id and name
	go hello()

	fmt.Println("end of the main")

	// worst case to wait for goroutines
	time.Sleep(time.Second) // sleep would put the main in the blocking state
	// go scheduler would pick up another goroutine if available

}

func hello() {
	fmt.Println("Hello, world!")
}
