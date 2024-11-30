package postgres

import (
	"fmt"
	"fundamentals/practice/day7/question3/app/stores"
	"log"
)

/*
In mysql.go create a Conn struct which stores db connection in string form

	Create three methods over Conn struct (Create(user) error, Update(name) error, Delete(id) error
	//Note:- return nil for the errors, assume no error would happen
	//For now add simple print statement without doing any actual work

	In postgres.go repeat the same steps as we did in mysql.go
*/

type PostgresConn struct {
	dbConnection string
}

func NewPostgresConnection(conn string) PostgresConn {
	if conn == "" {
		log.Fatal("empty connection string")
	}
	return PostgresConn{dbConnection: conn}
}

func (c PostgresConn) Create(u stores.User) error {
	fmt.Println("Creating user: ", u, " over connection", c)
	return nil
}

func (c PostgresConn) Update(u stores.User) error {
	fmt.Println("Updating user: ", u, " over connection", c)
	return nil
}

func (c PostgresConn) Delete(id int) error {
	fmt.Println("Deleting user with id: ", id, " over connection", c)
	return nil
}