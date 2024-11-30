package compositedatatypes

import (
	"fmt"
	"log"
)

// Conf - db configuration
type Conf struct {
	db string
}

// New functions are used to intialize struct with some config values,
// See proj-struct folder for proper implementation

// NewConf - creates and returns a new Conf object
func NewConf(conn string) Conf {
	if conn == "" {
		// avoid in production until and unless you want your app to stop working
		// this will crash the program
		log.Fatal("empty connection string")
	}
	// try to open the connection, and if it is successful, return the connection
	return Conf{db: conn}
}

// AddToDb - adding row to DB
func (c Conf) AddToDb() {
	fmt.Println("adding to db", c.db)
}

// TestNewConfFunctionality - Tests NewConf functionality in struct
func TestNewConfFunctionality() {
	c := NewConf("mysql://root:123456@localhost/test")
	c.AddToDb()
	log.Println(c.db)
}
