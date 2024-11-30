package main

import "fmt"

/*
Q3. Create a function that takes a list of users

	This func can append new values to the list or change the existing elems
	But Make sure this function can't modify the original slice
	that was created in the main function
*/

// User of type string
type User string

func main() {
	users := []User{"Komal", "Soni", "Shubhi", "Shruti"}
	var newUsers = AppendUpdateUsers(users)
	InspectUserSlice("users", users)
	InspectUserSlice("new_users", newUsers)
}

// AppendUpdateUsers - append on and update users list without changing original user list
func AppendUpdateUsers(users []User) []User {
	newUsers := make([]User, len(users))
	copy(newUsers, users)
	newUsers[0] = "SomeName"
	newUsers = append(newUsers, "AppenddedName")
	return newUsers
}

// InspectUserSlice - inspects user slice
func InspectUserSlice(name string, slice []User) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Printf("memory address %p\n", slice)
	fmt.Println(slice)
	fmt.Println()
}
