package main

import (
	"rest/config"
	"rest/database"
	"rest/users"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Load()
}

func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	v1 := router.Group("/v1")
	users.RegisterRoutes(v1)

	conf := config.New()

	database.Connect(conf.Database.Host, conf.Database.Name, conf.Database.User, conf.Database.Password)
	database.GetDB().AutoMigrate(&users.UserModel{})

	router.Run(":" + conf.PORT)

}
