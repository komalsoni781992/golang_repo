package stores

/*
In store.go

	Create an interface that postgres and mysql package can implement
*/
type DBConnection interface {
	Create(u User) (error, bool)
	Update(u User) (error, bool)
	Delete(id int) (error, bool)
}
