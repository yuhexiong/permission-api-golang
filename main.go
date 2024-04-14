package main

import (
	"fmt"
	"io"
	"log"
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
	// open logFile
	logFile, err := os.OpenFile("logFile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		util.RedLog("Unable to open file", err)
	}
	defer logFile.Close()

	// setup log
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

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
		util.RedLog("api port error:", err)
		panic(err)
	}

	util.GreenLog("Server run at port: %d", apiPort)
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
}
