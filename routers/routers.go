package routers

import (
	"database/sql"
	"miniProject/internals/handlers"
	"miniProject/internals/repositories"
	"miniProject/internals/services"
	"os/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, db *sql.DB) {
	//user auth
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)


	//repo init
	categoryRepo := repositories.NewCategoryRepository(db)
	//service init
	categoryService := services.NewCategoryService(categoryRepo)
	//handler init
	categoryHandler := handlers.NewCategoryHandler(categoryService)


	api := r.Group("/api")
	{
		// testing endpoint
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		
		userRoutes := api.Group("/users")
		{
			userRoutes.POST("/login", userHandler.Login)
		}

		categoryRoutes := api.Group("/categories")
		{
			// GET /api/categories
			categoryRoutes.GET("", categoryHandler.GetAll) 
			
			// POST /api/categories
			categoryRoutes.POST("", categoryHandler.Create) 
			
			// GET /api/categories/:id
			categoryRoutes.GET("/:id", categoryHandler.GetByID) 
			
			// DELETE /api/categories/:id
			categoryRoutes.DELETE("/:id", categoryHandler.Delete)
			
			// PUT /api/categories/:id
			categoryRoutes.PUT("/:id", categoryHandler.Update)
			
			// TODO: Tambahkan route Update (PUT/PATCH /api/categories/:id)
			// TODO: Tambahkan route Get Books by Category ID (GET /api/categories/:id/books)
		}
	}
}