package main

import (
 "github.com/moh-hosseini98/go-crud/initializers"
 "github.com/moh-hosseini98/go-crud/models"
)

func init() {
 initializers.LoadEnvVariables()
 initializers.ConnectDB()
}

func main() {
 initializers.DB.AutoMigrate(&models.Todo{})
}