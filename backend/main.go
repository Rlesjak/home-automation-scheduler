package main

import (
	appconf "rlesjak.com/ha-scheduler/config"
	Models "rlesjak.com/ha-scheduler/model"
)

func main() {

	config := appconf.GetAppConfig()

	// Connect to the database
	Models.ConnectDatabase(config)
}
