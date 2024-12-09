package main

import "fmt"

/*q2. Create an Interface with one method square(int)
  Create a type that implements this interface
  The type has one field that stores the currentSquare value
  Create a function Operation that can call the square method using the interface

  In the main function, create a nil pointer to the concrete type
  Pass the value to the operation function

  Operation function calls the method that implements the interface

  Try to do recovery from panic at different levels*/

func main() {
	var s *SquareImpl
	defer recoverPanic()
	callSquare(s)
	fmt.Println("end of main")
}

type SquareInterface interface {
	square(int)
}

type SquareImpl struct {
	squareVal int
}

func (s *SquareImpl) square(num int) {
	s.squareVal = num * num
}

func callSquare(s SquareInterface) {
	s.square(4)
	fmt.Println("end of callSquare")
}

func recoverPanic() {
	msg := recover()
	if msg != nil {
		fmt.Println("Panic happened:", msg)
		return
	}
}
