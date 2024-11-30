package user

import "fmt"

var userName string

// AddToDb - authenticates name and add to DB
func AddToDb(name string) string {
	userName = name
	fmt.Println(userName)
	//auth.Authenticate(name)
	return name
}
