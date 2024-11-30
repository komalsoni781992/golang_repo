package stores

/*
In store.go

	Create an interface that postgres and mysql package can implement
*/
type DBConnection interface {
	Create(u User) bool
	Update(u User) bool
	Delete(id int) bool
	//FetchAll()
	//FetchUser(id int) User
}
