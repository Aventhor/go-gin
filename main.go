package main

import (
	"rest/config"
	"rest/database"
	"rest/docs"
	"rest/users"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.Load()
}

// @title Swagger API
// @version 1.0
func main() {
	router := gin.Default()
	router.Use(gin.Logger())

	docs.SwaggerInfo.BasePath = "/v1"

	v1 := router.Group("/v1")
	users.RegisterRoutes(v1)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	conf := config.New()

	database.Connect(conf.Database.Host, conf.Database.Name, conf.Database.User, conf.Database.Password)
	database.GetDB().AutoMigrate(&users.UserModel{})

	router.Run(":" + conf.PORT)

}
