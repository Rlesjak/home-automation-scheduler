package appconf

import (
	"log"
	"os"

	"github.com/gookit/config/v2"
	"github.com/joho/godotenv"
)

type AppConfig struct {
	env_db_host string
	env_db_user string
	env_db_pwd  string
	env_db_name string
	env_db_port string
	name        string
}

func GetAppConfig() config.Config {
	if os.Getenv("DB_HOST") == "" {
		// If not, try sourcing them from .env.local file
		if err := godotenv.Load(".env.local"); err != nil {
			panic("Failed to load .env.local file to source enviroment variables!")
		}
	}

	cnf := config.New("app-config")

	cnf.LoadOSEnvs(map[string]string{
		"DB_HOST":     "env_db_host",
		"DB_USER":     "env_db_user",
		"DB_PASSWORD": "env_db_pwd",
		"DB_NAME":     "env_db_name",
		"DB_PORT":     "env_db_port",
		"MQ_HOST":     "env_mqtt_host",
		"MQ_USER":     "env_mqtt_user",
		"MQ_PWD":      "env_mqtt_pwd",
	})

	// Load configuration
	if err := cnf.LoadFiles("conf.json"); err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("config data: \n %#v\n", cnf.Data())

	// appConfig := AppConfig{}

	// if err := cnf.BindStruct("appConfig", &appConfig); err != nil {
	// 	log.Fatal("Failed to bind loaded config to AppConfig type!", err)
	// }

	return *cnf
}
