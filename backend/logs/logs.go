package logs

import (
	"log"
	"os"
)

var Info *log.Logger
var Error *log.Logger

func InitLoggers() {

	if err := os.MkdirAll("log", 0755); err != nil {
		log.Fatal(err)
	}

	// Create Info logger
	infoFile, err := os.OpenFile("log/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	Info = log.New(infoFile, "[INFO]", log.Ldate|log.Ltime)
	Info.Println("Initialised INFO logger.")

	// Create error logger
	errorFile, err := os.OpenFile("log/errors.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// Create a logger for errors
	Error = log.New(errorFile, "[ERROR]", log.Ldate|log.Ltime|log.Lshortfile)
	Info.Println("Initialised ERROR logger")
}

func CreateCustomLoggerFile(filename string) *os.File {
	// Create Info logger
	file, err := os.OpenFile("log/"+filename+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
