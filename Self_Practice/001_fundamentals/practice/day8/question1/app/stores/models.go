package stores

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Id    int
}

type Users map[int]*User

// Write implements io.Writer.
/*func (u User) Write(p []byte) (n int, err error) {
	u.Name += string(p)
	fmt.Println("user:", u.Name)
	return len(p), nil
}*/

func (u *User) Write(p []byte) (n int, err error) {
	u.Name += string(p)
	fmt.Println("user:", u.Name)
	return len(p), nil
}
