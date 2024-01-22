package main

import (
	"os"
	"permission-api/config"
	"permission-api/controller/permissionController"
	"permission-api/router"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := gin.Default()

	// database
	config.ConnectDB()

	// viper
	permissionController.InitViper()

	// Router
	router.UserRoute(app)

	apiPort := os.Getenv("API_PORT")
	app.Run(":" + apiPort)
}
