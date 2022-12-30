//This file will be used to run the migrations on our models

package main

import (
	"make5app-golang/initializers"
	"make5app-golang/models"
)

func init() {
	// This function will be called before the main function

	initializers.LoadEnvVariables()
	initializers.ConnnectToDB()
}

func main() {
	//mention the models here which need to be migrated
	initializers.DB.AutoMigrate(&models.User{})
}
