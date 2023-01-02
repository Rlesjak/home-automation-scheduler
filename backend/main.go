package main

import (
	appconf "rlesjak.com/ha-scheduler/config"
	"rlesjak.com/ha-scheduler/logs"
	models "rlesjak.com/ha-scheduler/model"
	"rlesjak.com/ha-scheduler/scheduler"
	"rlesjak.com/ha-scheduler/server"
)

func main() {

	// Initialise logger singletons
	logs.InitLoggers()

	// Create config
	config := appconf.GetAppConfig()

	// Connect to the database
	models.ConnectDatabase(config)

	// Initialise scheduler
	scheduler.InitScheduler(config)

	// Start listening for http requests
	server.StartServer(config)
}
