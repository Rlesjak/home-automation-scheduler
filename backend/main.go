package main

import (
	appconf "rlesjak.com/ha-scheduler/config"
	Models "rlesjak.com/ha-scheduler/model"
	"rlesjak.com/ha-scheduler/server"
)

func main() {

	config := appconf.GetAppConfig()

	// Connect to the database
	Models.ConnectDatabase(config)

	// Start listening for http requests
	server.StartServer(config)
}
