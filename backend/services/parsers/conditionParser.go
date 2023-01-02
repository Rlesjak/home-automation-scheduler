package parsers

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

var DaysOfWeek []string = []string{
	"monday",
	"tuesday",
	"wednesday",
	"thursday",
	"friday",
	"saturday",
	"sunday",
}

type ParsedCondition struct {
	EveryNof int
	At       sql.NullString
	Events   []string
}

func condParserError(msg string) error {
	return errors.New("Condition parser error: " + msg)
}

func ParseCondition(conditionString string) (ParsedCondition, error) {
	startsWith := "every "
	statementDelimiter := " "
	eventsDelimiter := ","

	var parsedCondition ParsedCondition

	// Check if condition string starts with correct pattern
	if !strings.HasPrefix(conditionString, startsWith) {
		return parsedCondition, condParserError("Condition string must start with '" + startsWith + "'")
	}

	// Trim starting pattern
	withNoPrefix := strings.TrimPrefix(conditionString, startsWith)

	// Split string in to statements
	// ex. "3 monday,tuseday,friday at 12:30" --> ["3", "monday,tuseday,friday", "at", "12:30"]
	conditionStringParts := strings.Split(withNoPrefix, statementDelimiter)

	// Convert string to integer
	everyNof, err := strconv.Atoi(conditionStringParts[0])
	if err != nil {
		return parsedCondition, condParserError(err.Error())
	}
	// Check if integer is in allowed range
	if everyNof < 0 || everyNof > 999 {
		return parsedCondition, condParserError("In 'every n ...', n must be between 1 and 999")
	}
	parsedCondition.EveryNof = everyNof

	// If there are 2 statements condition dos not have "at" statement
	// If there are 4 statements condition dos have "at" statement
	if len(conditionStringParts) == 4 {
		if conditionStringParts[2] != "at" {
			return parsedCondition, condParserError("The third statement must be 'at'")
		}

		parsedCondition.At.String = conditionStringParts[3]
		parsedCondition.At.Valid = true
	}

	//TODO: Check if only weekdays

	// Split delimited events in to array of strings
	parsedCondition.Events = strings.Split(conditionStringParts[1], eventsDelimiter)

	return parsedCondition, nil
}
