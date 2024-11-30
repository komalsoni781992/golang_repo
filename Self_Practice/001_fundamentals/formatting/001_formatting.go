package formatting

import "fmt"

// BasicFormatting - formatting with printf
func BasicFormatting() {
	//use of Print
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3783)
	fmt.Print(" lbs, and I would be ")
	fmt.Print(41 * 365 / 687)
	fmt.Print(" years old.\n")

	//use of Println
	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be", 41*365.2425/687, "years old.")

	//use of Printf
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)

	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)
	fmt.Printf("My weight on the surface of %s is %f lbs.\n", "Earth", 149.0)
	//fmt.Printf("My weight on the surface of %f is %s lbs.\n", "Earth", 149.0) - compile time error
	fmt.Printf("My weight on the surface of %#v is %#v lbs.\n", "Earth", 149.0)
	fmt.Printf("My weight on the surface of %T is %T lbs.\n", "Earth", 149.0)
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
