package main

import (
	"make5app-golang/initializers"

	"github.com/gin-gonic/gin"

	"make5app-golang/controllers"
)

func Landing(c *gin.Context) {
	c.IndentedJSON(200, gin.H{
		"Function":         "API Endpoint",
		"To post a user":   "/user",
		"To get all users": "/user",
	})
}

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/", Landing)

	r.POST("/user", controllers.PostUser)
	r.GET("/user", controllers.GetUsers)
	r.Run()
}
