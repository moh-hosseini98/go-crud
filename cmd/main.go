package main

import (
 "github.com/gin-gonic/gin"
 "github.com/moh-hosseini98/go-crud/initializers"
 "github.com/moh-hosseini98/go-crud/routes"
)

func init() {
 initializers.LoadEnvVariables()
 initializers.ConnectDB()
}

func main() {

 r := gin.Default()

 // Todo Routes
 routes.TodoRoutes(r)

 r.Run()
}