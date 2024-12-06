package main

import (
	"fmt"
	"fundamentals/practice/day8/question1/app/stores"
	"fundamentals/practice/day8/question1/app/stores/mysql"
	"fundamentals/practice/day8/question1/app/stores/postgres"
)

/*Call postgres and mysql package methods using interface variable*/

func main() {
	mysqlConnection := mysql.NewMySQLConnection("Mysql")
	postgresConnection := postgres.NewMyPostgresConnection("Postgres")
	TestDBMethods(mysqlConnection)
	TestDBMethods(postgresConnection)
}

// TestDBMethods - tests db methods
func TestDBMethods(r stores.DBConnection) {
	user := stores.User{Name: "Komal", Email: "soni.komal@gmail.com", ID: 7}
	user2 := stores.User{Name: "Soni", Email: "soni.komal@gmail.com", ID: 1}
	if result := r.Create(user); result {
		fmt.Println("User created successfully.")
	}
	if result := r.Create(user2); result {
		fmt.Println("User created successfully.")
	}
	if result := r.Update(user); result {
		fmt.Println("User updated successfully.")
	}
	if result := r.Delete(1); result {
		fmt.Println("User deleted successfully.")
	}
	if f, ok := r.(mysql.MySQLConn); ok {
		p := f.FetchUser(7)
		fmt.Println(p)
		fmt.Println("Printing all usesr:")
		f.FetchAll()
	}

	if f, ok := r.(postgres.PostgresConn); ok {
		p := f.FetchUser(7)
		fmt.Println(p)
		fmt.Println("Printing all usesr:")
		f.FetchAll()
	}
}
