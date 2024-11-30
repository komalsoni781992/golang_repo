package main

import (
	"fmt"
)

/*Q2. Create an initial list of users.
  Create a new slice named as emp from a part users slice.
  Check both slices length and capacity.

  Append some more user names to the *user* slice.
  Check len, cap again for both slices
  Print the values as well

  Append some emp names to the *emp* slice
  Check len, cap again for both slices
  Print the values as well

  Compare the contents and capacities of the original list and the modified slice.
  Try to understand and visualize what is happening*/

func main() {
	// Solution - With memory sharing
	//Create an initial list of users.
	usersOne := []string{"Komal", "Soni", "Shubhi", "Shruti"}
	InspectStringSlice("usersOne", usersOne)

	//Create a new slice named as emp from a part users slice.
	empOne := usersOne[2:]
	fmt.Println("emp created by make")
	InspectStringSlice("usersOne", usersOne)
	InspectStringSlice("empOne", empOne)

	//Append some more user names to the *user* slice.
	//Check len, cap again for both slices
	//Print the values as well
	fmt.Println("New usesr added to usesr")
	usersOne = append(usersOne, "Hello", "World")
	InspectStringSlice("usersOne", usersOne)
	InspectStringSlice("empOne", empOne)

	//Append some emp names to the *emp* slice
	//Check len, cap again for both slices
	//Print the values as well
	fmt.Println("New users added to emp")
	empOne = append(empOne, "Bye", "Bye")
	InspectStringSlice("usersOne", usersOne)
	InspectStringSlice("empOne", empOne)

	//Solution -  Without memory sharing

	//Create an initial list of users.
	users := []string{"Komal", "Soni", "Shubhi", "Shruti"}
	InspectStringSlice("users", users)

	//Create a new slice named as emp from a part users slice.
	emp := make([]string, len(users[2:]))
	fmt.Println("emp created by make")
	InspectStringSlice("users", users)
	InspectStringSlice("emp", emp)

	copy(emp, users[2:])
	fmt.Println("some users copied to emp")
	InspectStringSlice("emp", emp)

	//Append some more user names to the *user* slice.
	//Check len, cap again for both slices
	//Print the values as well
	fmt.Println("New usesr added to usesr")
	users = append(users, "Hello", "World")
	InspectStringSlice("users", users)
	InspectStringSlice("emp", emp)

	//Append some emp names to the *emp* slice
	//Check len, cap again for both slices
	//Print the values as well
	fmt.Println("New usesr added to emp")
	emp = append(emp, "Bye", "Bye")
	InspectStringSlice("usesr", users)
	InspectStringSlice("emp", emp)
}

// InspectStringSlice - inspects underlying array for length and capacity and memory address
func InspectStringSlice(name string, slice []string) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()
}
