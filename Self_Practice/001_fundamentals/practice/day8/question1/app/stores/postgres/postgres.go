package postgres

import (
	"fmt"
	"fundamentals/practice/day8/question1/app/stores"
	"log"
)

/*
In mysql.go create a Conn struct which stores db connection in string form

	Create three methods over Conn struct (Create(user) error, Update(name) error, Delete(id) error
	//Note:- return nil for the errors, assume no error would happen
	//For now add simple print statement without doing any actual work

	In postgres.go repeat the same steps as we did in mysql.go
*/

// PostgresConn - struct for postgres connection
type PostgresConn struct {
	dbConnection string
	postgresMap  stores.Users
}

// NewMyPostgresConnection - creates new postgres connection
func NewMyPostgresConnection(conn string) PostgresConn {
	if conn == "" {
		log.Fatal("empty connection string")
	}
	return PostgresConn{dbConnection: conn, postgresMap: make(stores.Users, 1)}
}

// Create - creates user
func (c PostgresConn) Create(u stores.User) bool {
	fmt.Println("Creating user: ", u, " over connection", c)
	_, ok := c.postgresMap[u.ID] // user, ok (true,false) if that value exist or not
	if !ok {
		c.postgresMap[u.ID] = &u
		return true
	}
	return false
}

// Update - updates user
func (c PostgresConn) Update(u stores.User) bool {
	fmt.Println("Updating user: ", u, " over connection", c)
	_, ok := c.postgresMap[u.ID] // user, ok (true,false) if that value exist or not
	if ok {
		c.postgresMap[u.ID] = &u
		return true
	}
	return false
}

// Delete - deletes user
func (c PostgresConn) Delete(id int) bool {
	fmt.Println("Deleting user with id: ", id, " over connection", c)
	_, ok := c.postgresMap[id] // user, ok (true,false) if that value exist or not
	if ok {
		delete(c.postgresMap, id)
		return true
	}
	return false
}

// FetchAll - fetches all users
func (c PostgresConn) FetchAll() {
	for key, value := range c.postgresMap {
		fmt.Println(key, value)
	}
}

// FetchUser - fetch usr by ID
func (c PostgresConn) FetchUser(id int) *stores.User {
	user, ok := c.postgresMap[id]
	if ok {
		return user
	}
	return nil
}
