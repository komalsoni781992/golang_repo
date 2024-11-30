package main

import (
	"fmt"
	"fundamentals/practice/day7/question3/app/stores"
	"fundamentals/practice/day7/question3/app/stores/mysql"
	"fundamentals/practice/day7/question3/app/stores/postgres"
)

func main() {
	mysqlConnection := mysql.NewMySqlConnection("MySql")
	postgresConnection := postgres.NewPostgresConnection("Postgres")
	TestDBMethods((mysqlConnection))
	TestDBMethods((postgresConnection))
}

func TestDBMethods(r stores.DBConnection) {
	user := stores.User{Name: "Komal", Email: "soni.komal@gmail.com", Id: 7}
	user2 := stores.User{Name: "Soni", Email: "soni.komal@gmail.com", Id: 1}
	err := r.Create(user)
	if err == nil {
		fmt.Println("User created successfully.")
	}
	err = r.Create(user2)
	if err == nil {
		fmt.Println("User created successfully.")
	}
	err = r.Update(user)
	if err == nil {
		fmt.Println("User updated successfully.")
	}
	err = r.Delete(1)
	if err == nil {
		fmt.Println("User deleted successfully.")
	}
}
