package users

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup) {
	r := rg.Group("/users")

	r.GET("/", UserList)
	r.GET("/:id", UserRetrieve)
	r.POST("/", UserCreate)
	r.PATCH("/:id", UserUpdate)
	r.DELETE("/:id", UserDelete)
}

func UserList(c *gin.Context) {
	u, err := FindAll(nil)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return
	}

	serializer := UsersSerializer{c, u}
	c.JSON(http.StatusOK, serializer.Response())
}

func UserRetrieve(c *gin.Context) {
	id := c.Param("id")
	u, err := FindOne(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return
	}

	serializer := UserSerializer{c, u}
	c.JSON(http.StatusOK, serializer.Response())
}

func UserCreate(c *gin.Context) {
	var user UserModel
	c.BindJSON(&user)

	fmt.Print(user)

	err := CreateOne(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad")
		return
	}
	serializer := UserSerializer{c, user}
	c.JSON(http.StatusCreated, gin.H{"user": serializer.Response()})
}

func UserUpdate(c *gin.Context) {
	id := c.Param("id")
	userModel, err := FindOne(id)

	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return
	}

	c.BindJSON(&userModel)

	if err := userModel.Update(userModel); err != nil {
		c.JSON(http.StatusBadRequest, "bad")
		return
	}

	serializer := UserSerializer{c, userModel}
	c.JSON(http.StatusOK, serializer.Response())
}

func UserDelete(c *gin.Context) {
	id := c.Param("id")

	userModel, err := FindOne(id)
	if err != nil {
		c.JSON(http.StatusNotFound, "not found")
		return
	}

	c.BindJSON(&userModel)

	if err := userModel.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, "error")
		return
	}

	serializer := UserSerializer{c, userModel}
	c.JSON(http.StatusOK, serializer.Response())
}
