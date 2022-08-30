package controllers

import (
	"log"
	"net/http"
	"todolist/initiliazers"
	"todolist/models"

	"github.com/gin-gonic/gin"
)


// ------------------------------------post method---------------------------------------------------//

func PostsCreate(c *gin.Context) {

	// giving struct same as our database to send our data 

	var body struct{ 
		Tname string         
		Tdescription string 
	}

	// bind will bind the data in the struct 

	  c.Bind(&body)

	  // Taskname and task description are from our main model....in this we wll pass our values via body struct which is same as our database
 
	post := models.Post{TaskName: body.Tname, TaskDescription: body.Tdescription}

	// herr we are creating our post which we send at its address

    result := initiliazers.DB.Create(&post) 

	 if result.Error != nil {
		log.Println("error faced")
	 }

	c.JSON(http.StatusOK, gin.H{
		"posts": post,
	})
}

//------------------------------------------------ get method --------------------------------------------------//

func PostsIndex(c *gin.Context) {
	var posts[] models.Post

    initiliazers.DB.Find(&posts)

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

//----------------------------------Get one request -----------------------------------//

func PostShow(c *gin.Context) {

	// it willbe like localhost :8080/posts/id

	id := c.Param("id")
	var post models.Post
    initiliazers.DB.First(&post, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

//----------------------------------updating  request -----------------------------------//

func PostUpdate(c *gin.Context){

	// selection the id 

	id := c.Param("id")

	// giving the post structure 

	var body struct{ 
		Tname string         
		TDescription string 
	}

	  c.Bind(&body)

	  // selecting particular post by this 

	  var post models.Post
	  initiliazers.DB.First(&post, id)

	  // updating our value

	  initiliazers.DB.Model(&post).Updates(models.Post{TaskName: body.Tname , TaskDescription: body.TDescription,})

	  // show the result

	  c.JSON(http.StatusOK, gin.H{
		"Update": post,
	  })

}

// ----------------  deleting request ---------------------------//

func DeleteReq(c *gin.Context) {

	// selecting the id 

	id := c.Param("id")

	// selection of the row 

	initiliazers.DB.Delete(&models.Post{},id)

	// it will delete the data for a particular id 
}