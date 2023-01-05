package parsers

import (
	"errors"
	"strings"

	"rlesjak.com/ha-scheduler/integrations"
	"rlesjak.com/ha-scheduler/scheduler"
)

func commandParserError(msg string) error {
	return errors.New("Command parser error: " + msg)
}

func ParseCommand(rawCommand string) (scheduler.JobCommand, error) {
	parsedCommand := scheduler.JobCommand{}

	// Separate integration from command
	splitString := strings.Split(rawCommand, ":")
	if len(splitString) != 2 {
		return parsedCommand, commandParserError("Incorrect command structure, must be '[integration]:[command]'")
	}

	// Get handler struct for the given integration
	integrationHandler, intFuncErr := integrations.GetIntegrationHandler(splitString[0])
	if intFuncErr != nil {
		return parsedCommand, commandParserError(intFuncErr.Error())
	}

	// Validate provided command with the provided integration
	if err := integrationHandler.Validate(splitString[1]); err != nil {
		return parsedCommand, commandParserError(err.Error())
	}

	// If validation checks out, return jobCommand struct
	// With a call to the integration handler function
	parsedCommand.Run = func() {
		integrationHandler.Handler(splitString[1])
	}

	return parsedCommand, nil
}
