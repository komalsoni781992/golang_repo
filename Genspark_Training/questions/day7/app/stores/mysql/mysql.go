package mysql

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

type MySqlConn struct {
	dbConnection string
	mysqlMap     users
}

func NewMySqlConnection(conn string) MySqlConn {
	if conn == "" {
		log.Fatal("empty connection string")
	}
	return MySqlConn{dbConnection: conn, mysqlMap: make(users, 1)}
}

func (c MySqlConn) Create(u stores.User) (error, bool) {
	fmt.Println("Creating user: ", u, " over connection", c)
	_, ok := c.mysqlMap[u.Id] // user, ok (true,false) if that value exist or not
	if !ok {
		c.mysqlMap[u.Id] = &u
		return nil, true
	}
	return nil, false
}

func (c MySqlConn) Update(u stores.User) (error, bool) {
	fmt.Println("Updating user: ", u, " over connection", c)
	_, ok := c.mysqlMap[u.Id] // user, ok (true,false) if that value exist or not
	if ok {
		c.mysqlMap[u.Id] = &u
		return nil, true
	}
	return nil, false
}

func (c MySqlConn) Delete(id int) (error, bool) {
	fmt.Println("Deleting user with id: ", id, " over connection", c)
	_, ok := c.mysqlMap[id] // user, ok (true,false) if that value exist or not
	if ok {
		delete(c.mysqlMap, id)
		return nil, true
	}
	return nil, false
}

func (c MySqlConn) FetchAll() {
	for key, value := range c.mysqlMap {
		fmt.Println(key, value)
	}
}

func (c MySqlConn) FetchUser(id int) {
	user, ok := c.mysqlMap[id]
	if ok {
		fmt.Println(user)
	}
}
