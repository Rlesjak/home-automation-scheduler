package parsers

import "rlesjak.com/ha-scheduler/scheduler"

func ParseJob(condition string, command string) (ParsedCondition, scheduler.JobCommand, error) {
	// Parse given condition
	parsedCondition, conditionParseErr := ParseCondition(condition)
	if conditionParseErr != nil {
		return ParsedCondition{}, scheduler.JobCommand{}, conditionParseErr
	}

	// Parse given command string
	parsedCommand, commandParseErr := ParseCommand(command)
	if commandParseErr != nil {
		return ParsedCondition{}, scheduler.JobCommand{}, commandParseErr
	}

	return parsedCondition, parsedCommand, nil
}
