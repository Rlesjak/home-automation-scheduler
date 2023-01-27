package main

import (
	"os"

	appconf "rlesjak.com/ha-scheduler/config"
	"rlesjak.com/ha-scheduler/integrations"
	"rlesjak.com/ha-scheduler/logs"
	models "rlesjak.com/ha-scheduler/model"
	"rlesjak.com/ha-scheduler/scheduler"
	"rlesjak.com/ha-scheduler/server"
	"rlesjak.com/ha-scheduler/services"
)

func main() {

	// Initialise logger singletons
	logs.InitLoggers()

	// Create config
	config := appconf.GetAppConfig()

	// Connect to the database
	dbErr := models.ConnectDatabase(config)
	if dbErr != nil {
		// If connection to the database cannot be established
		// server should not start
		os.Exit(1)
		return
	}

	// Initialise scheduler
	scheduler.InitScheduler(config)

	// Initialise integrations
	integrations.InitialiseIntegrations(config)

	// Initialise services
	srvcsErr := services.InitialiseServices(config)
	if srvcsErr != nil {
		// If services failed initialisation server should not start
		os.Exit(1)
		return
	}

	// Start listening for http requests
	server.StartServer(config)
}
