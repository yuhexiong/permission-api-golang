package main

import (
	"fmt"
	"net/http"
	"os"
	"permission-api/config"
	"permission-api/controller"
	"permission-api/controller/permissionController"
	"permission-api/router"
	"permission-api/util"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	// database
	config.ConnectDB()

	// viper
	permissionController.InitViper()

	// Router
	router := router.InitRouter()

	// init AdminUser
	controller.InitAdminUser()

	// setup http server
	apiPortStr := os.Getenv("API_PORT")
	apiPort, err := strconv.Atoi(apiPortStr)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", apiPort),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := server.ListenAndServe(); err != nil {
		util.RedLog("Server error:", err)
	}
	util.GreenLog("Server run at port:", apiPort)
}
