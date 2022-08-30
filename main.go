package main

import (
	"todolist/controllers"
	"todolist/initiliazers"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func init() {
	initiliazers.LoadEnvVariables()
	initiliazers.ConnectDB()
}
func main() {
	//---------------------- initilizing our db -------------------------//
	initiliazers.DB.AutoMigrate(&models.Post{})

	r := gin.Default()
	//---------------------our all requests ----------------------------//

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)
	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)
	r.DELETE("/posts/:id", controllers.DeleteReq)
	r.Run()
}