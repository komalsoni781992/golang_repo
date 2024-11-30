package postgres

import (
	"fmt"
	"interfacesintoduction/app/stores"
	"log"
)

/*
In mysql.go create a Conn struct which stores db connection in string form

	Create three methods over Conn struct (Create(user) error, Update(name) error, Delete(id) error
	//Note:- return nil for the errors, assume no error would happen
	//For now add simple print statement without doing any actual work

	In postgres.go repeat the same steps as we did in mysql.go
*/
type users map[int]*stores.User

type PostgresConn struct {
	dbConnection string
	postgresMap  users
}

func NewMyPostgresConnection(conn string) PostgresConn {
	if conn == "" {
		log.Fatal("empty connection string")
	}
	return PostgresConn{dbConnection: conn, postgresMap: make(users, 1)}
}

func (c PostgresConn) Create(u stores.User) (error, bool) {
	fmt.Println("Creating user: ", u, " over connection", c)
	_, ok := c.postgresMap[u.Id] // user, ok (true,false) if that value exist or not
	if !ok {
		c.postgresMap[u.Id] = &u
		return nil, true
	}
	return nil, false
}

func (c PostgresConn) Update(u stores.User) (error, bool) {
	fmt.Println("Updating user: ", u, " over connection", c)
	_, ok := c.postgresMap[u.Id] // user, ok (true,false) if that value exist or not
	if ok {
		c.postgresMap[u.Id] = &u
		return nil, true
	}
	return nil, false
}

func (c PostgresConn) Delete(id int) (error, bool) {
	fmt.Println("Deleting user with id: ", id, " over connection", c)
	_, ok := c.postgresMap[id] // user, ok (true,false) if that value exist or not
	if ok {
		delete(c.postgresMap, id)
		return nil, true
	}
	return nil, false
}

func (c PostgresConn) FetchAll() {
	for key, value := range c.postgresMap {
		fmt.Println(key, value)
	}
}

func (c PostgresConn) FetchUser(id int) {
	_, ok := c.postgresMap[id]
	if ok {
		fmt.Println(c.postgresMap[id])
	}
}
