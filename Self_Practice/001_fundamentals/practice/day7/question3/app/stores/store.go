package stores

/*
In store.go

	Create an interface that postgres and mysql package can implement
*/
type DBConnection interface {
	Create(u User) error
	Update(u User) error
	Delete(id int) error
}
