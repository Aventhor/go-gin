package users

type CreateUserDto struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type UpdateUserDto struct {
	CreateUserDto
}
