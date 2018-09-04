package users

// UserService interfaces defines the list of methods that the service needs to provide
// for the user handlers
type Service interface {
	GetByID(ID string) (User, error)
	GetByEmail(Email string) (User, error)
	Create(*User) (User, error)
	Update(*User) (User, error)
}
