//controllers contain the main logic of the API's

package controllers

import (
	"make5app-golang/initializers"
	"make5app-golang/models"

	"github.com/gin-gonic/gin"
)

func PostUser(c *gin.Context) {
	//logic to create a user

	// this is like request.body in django
	var body struct {
		Name       string `json:"name"`
		RollNumber int    `json:"roll_number"`
		Email      string `json:"email"`
	}

	c.Bind(&body)

	user := models.User{Name: body.Name, RollNumber: body.RollNumber, Email: body.Email}

	result := initializers.DB.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		panic("Failed to create user!")
	}

	//message to return successfully user created
	c.IndentedJSON(200, gin.H{
		"message": "User created successfully!",
		"User":    user,
	})
}

func GetUsers(c *gin.Context) {

	// Get all records
	var users []models.User
	result := initializers.DB.Find(&users)
	// SELECT * FROM users;

	if result.Error != nil {
		panic("Failed to get users!")
	}

	c.IndentedJSON(200, gin.H{
		"message": "Users fetched successfully!",
		"Users":   users,
	})
}
