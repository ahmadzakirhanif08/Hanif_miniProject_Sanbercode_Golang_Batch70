package miniproject

import (
	"fmt"
	"log"
	"miniProject/pkg/configs"
	"miniProject/routers"


	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	//check .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading env")
	}

	//init database
	db := configs.InitDB()
	defer configs.CloseDB()

	//run migration
	configs.RunMigrations(db)

	//init gin engine
	r := gin.Default()

	//routers setup
	routers.SetupRouter(r, db)

	//run server
	port := "8080"
	fmt.Printf("Server running on http://localhost:%s\n", port)
	r.Run(":" + port)
}