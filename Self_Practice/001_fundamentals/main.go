package main

import (
	"fmt"
	"fundamentals/controlstructures"
	"time"
)

func main() {
	//basicdatatypes.TestString()
	fmt.Println(controlstructures.LeapYear(2004))
	time.Sleep(time.Second)
	fmt.Println(controlstructures.LeapYear(3000))
}
