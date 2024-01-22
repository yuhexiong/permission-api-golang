package main

import (
	"os"
	"permission-api/config"
	"permission-api/router"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := gin.Default()

	// Run database
	config.ConnectDB()

	// Router
	router.UserRoute(app)

	apiPort := os.Getenv("API_PORT")
	app.Run(":" + apiPort)
}
