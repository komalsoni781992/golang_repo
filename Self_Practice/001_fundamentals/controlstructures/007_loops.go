package controlstructures

import (
	"fmt"
	"math"
)

// TestForLoop - 4 types of for loop
func TestForLoop() {
	//var is not legal here.
	//Just like variable declarations in if statements, you can shadow a variable here.

	//complete for
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//condition only for
	i := 1
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}

	// infinite loop
	/*for {
		fmt.Println("Hello")
	}*/

	//do-while loop
	for {
		// things to do in the loop
		var x int
		if _, err := fmt.Scanln(&x); err == nil && x == 0 {
			fmt.Println(x)
			break
		}
	}

	// if else makes code confusing
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// continue simplifies above code
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}

	// for range
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}

	// return value can be ignored with underscore
	evenVals = []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		fmt.Println(v)
	}

	// second value can be ignored as it is
	evenVals = []int{2, 4, 6, 8, 10, 12}
	for i := range evenVals {
		fmt.Println(i)
	}

	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}

	// order of map scanning might every time:
	// In earlier Go versions, the iteration order for keys in a map was usually (but not always) the same if you inserted the same items into a map. This caused two problems.
	// People would write code that assumed that the order was fixed, and this would break at weird times.
	// If maps always hash items to the exact same values, and you know that a server is storing some user data in a map, you can actually slow down a server with an attack called Hash DoS
	// by sending it specially crafted data where all of the keys hash to the same bucket.
	m := map[string]int{
		"a": 1,
		"c": 3,
		"b": 2,
	}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}

	//There is one exception to this rule. In order to make it easier to debug and log maps, the formatting functions (like fmt.Println) always output maps with their keys in
	// ascending sorted order.

	// for range on string iterates over rune and not byte
	//It iterates over the runes, not the bytes.
	// Whenever a for-range loop encounters a multi-byte rune in a string, it converts the UTF-8 representation into a single 32-bit number and assigns it to the value.
	// The offset is incremented by the number of bytes in the rune.
	// If the for-range loop encounters a byte that doesn’t represent a valid UTF-8 value, the Unicode “Replacement Character” (hex value 0xfffd) is returned instead.
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
		}
		fmt.Println()
	}
	/*
		0 104 h
		1 101 e
		2 108 l
		3 108 l
		4 111 o

		0 97 a
		1 112 p
		2 112 p
		3 108 l
		4 101 e
		5 95 _
		6 960 π
		8 33 !
	*/

	// value is a copy so updation cannot be done like this
	evenVals = []int{2, 4, 6, 8, 10, 12}
	for _, v := range evenVals {
		//v *= 2
		fmt.Println(v)
	}
	fmt.Println(evenVals)

	// break and continue applies to for-range
	// if you launch goroutines in a for-range loop, you need to be very careful in how you pass the index and value to the goroutines

	samples = []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer
			}
		}
		fmt.Println()
	}
	/*
		0 104 h
		1 101 e
		2 108 l
		0 97 a
		1 112 p
		2 112 p
		3 108 l
	*/

	// confusing code
	evenVals = []int{2, 4, 6, 8, 10}
	for i, v := range evenVals {
		if i == 0 {
			continue
		}
		if i == len(evenVals)-2 {
			break
		}
		fmt.Println(i, v)
	}

	// simplified code
	evenVals = []int{2, 4, 6, 8, 10}
	for i := 1; i < len(evenVals)-1; i++ {
		fmt.Println(i, evenVals[i])
	}

	//This pattern does not work for skipping over the beginning of a string.
	//Remember, a standard for loop doesn’t properly handle multi-byte characters.
	//If you want to skip over some of the runes in a string, you need to use a for-range loop so that it will properly process runes for you.
}

/*
2. Write a program that prints out all the numbers between 1 and 100 that are
evenly divisible by 3 (i.e., 3, 6, 9, etc.).
*/

// PrintNumbersDivisibleBy3 program that prints out all the numbers between 1 and 100 that are evenly divisible by 3 (i.e., 3, 6, 9, etc.).
func PrintNumbersDivisibleBy3() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

// PrintFizzBuzz - Write a program that prints the numbers from 1 to 100, but for multiples of three, print “Fizz” instead of the number, and for the multiples of five, print “Buzz.”
// For numbers that are multiples of both three and five, print “FizzBuzz.
func PrintFizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
			continue
		} else if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		} else if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		fmt.Println(i)
	}
}

/*Write a program that finds the smallest number in this list:
x := []int{
	48,96,86,68,
	57,82,63,70,
	37,34,83,27,
	19,97, 9,17,
   }*/

// FindSmallestNumber - finds smallest element in list
func FindSmallestNumber() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	smallest := math.MaxInt
	for _, v := range x {
		if v < smallest {
			smallest = v
		}
	}
	fmt.Println(smallest)
}
