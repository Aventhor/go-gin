package users

import "github.com/gin-gonic/gin"

type UsersSerializer struct {
	C     *gin.Context
	Users []UserModel
}

type UserSerializer struct {
	C *gin.Context
	UserModel
}

type UserResponse struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
}

func (s *UsersSerializer) Response() []UserResponse {
	response := []UserResponse{}
	for _, user := range s.Users {
		serializer := UserSerializer{s.C, user}
		response = append(response, serializer.Response())
	}
	return response
}

func (s *UserSerializer) Response() UserResponse {
	response := UserResponse{ID: s.ID, Email: s.Email, FullName: s.GetName()}
	return response
}
