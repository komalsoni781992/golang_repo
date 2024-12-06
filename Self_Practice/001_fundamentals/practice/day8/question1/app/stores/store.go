package stores

/*
In store.go

	Create an interface that postgres and mysql package can implement
*/

// DBConnection - interface for DB activities
type DBConnection interface {
	Create(u User) bool
	Update(u User) bool
	Delete(id int) bool
	//FetchAll()
	//FetchUser(id int) User
}
