package logah

import (
	"log"
	"os"
)

var (
	Logger *log.Logger
)

// Initialize the logger
func init() {
	// Open or create the log file
	logFile, err := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create a new logger instance
	Logger = log.New(logFile, "LOG: ", log.Ldate|log.Ltime|log.Lshortfile)
}