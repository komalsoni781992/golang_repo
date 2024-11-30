package main

import (
	"fmt"
	"interfacesintoduction/app/stores"
	"interfacesintoduction/app/stores/mysql"
	"log"
)

/*Call postgres and mysql package methods using interface variable*/

func main() {
	/*mysqlConnection := mysql.NewMySqlConnection("Mysql")
	postgresConnection := postgres.NewMyPostgresConnection("Postgres")
	TestDBMethods((mysqlConnection))
	TestDBMethods((postgresConnection))*/
	user := &stores.User{Name: "Komal", Email: "soni.komal@gmail.com", Id: 7}
	//l := log.New(os.Stdout, "log: ", log.LstdFlags)
	l := log.New(user, "log: ", log.LstdFlags)
	l.Println("Hello, log")
}

func TestDBMethods(r stores.DBConnection) {
	user := stores.User{Name: "Komal", Email: "soni.komal@gmail.com", Id: 7}
	user2 := stores.User{Name: "Soni", Email: "soni.komal@gmail.com", Id: 1}
	err, result := r.Create(user)
	if err == nil && result {
		fmt.Println("User created successfully.")
	}
	err, result = r.Create(user2)
	if err == nil && result {
		fmt.Println("User created successfully.")
	}
	err, result = r.Update(user)
	if err == nil && result {
		fmt.Println("User updated successfully.")
	}
	err, result = r.Delete(1)
	if err == nil && result {
		fmt.Println("User deleted successfully.")
	}
	f, ok := r.(mysql.MySqlConn)
	if ok {
		f.FetchAll()
	}
}
