package integrations

import (
	"errors"

	"github.com/gookit/config/v2"
)

type IntegrationCommand struct {
	Handler  func(string)
	Validate func(string) error
}

func InitialiseIntegrations(config config.Config) {

	// Initialise/Connect mqtt
	initialiseMqttIntegration(config)
}

func GetIntegrationHandler(integrationName string) (IntegrationCommand, error) {

	switch integrationName {
	case "mqtt":
		return IntegrationCommand{
			Handler:  mqttIntegrationHanlder,
			Validate: mqttValidateCommand,
		}, nil
	}

	return IntegrationCommand{}, errors.New("Unknown integration '" + integrationName + "'")
}
