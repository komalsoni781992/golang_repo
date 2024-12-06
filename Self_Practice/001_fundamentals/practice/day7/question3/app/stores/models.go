package stores

import "fmt"

// User - struct with name, email and id
type User struct {
	Name  string
	Email string
	ID    int
}

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
