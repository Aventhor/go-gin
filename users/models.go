package users

import (
	"fmt"

	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Email     string
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *UserModel) GetName() string {
	fmt.Println(u.LastName, u.FirstName)
	return u.LastName + " " + u.FirstName
}
