package routes

import (
	"miniProject/controllers"
	"os"
    "log"
	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine { 
    user := os.Getenv("BASIC_AUTH_USER")
    password := os.Getenv("BASIC_AUTH_PASSWORD")
    if user == "" || password == "" {
        log.Println("FATAL ERROR: BASIC_AUTH_USER or BASIC_AUTH_PASSWORD is not set!") 

        user = "admin"
        password = "admin"
    }
    router := gin.Default()

    authorized := router.Group("/api")
    
    authorized.Use(gin.BasicAuth(gin.Accounts{
        user:    password,       
    }))
    {
        authorized.GET("/categories", controllers.FindCategories) 
        authorized.POST("/categories", controllers.CreateCategory)
        authorized.GET("/categories/:id", controllers.GetCategoryByID) 
        authorized.DELETE("/categories/:id", controllers.DeleteCategory)
        authorized.GET("/categories/:id/books", controllers.FindBooksByCategory)

        authorized.GET("/books", controllers.FindBooks)
        authorized.POST("/books", controllers.CreateBook)
        authorized.GET("/books/:id", controllers.GetBookByID)
        authorized.DELETE("/books/:id", controllers.DeleteBook)
    }

    return router
}