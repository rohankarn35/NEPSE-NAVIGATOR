package applog

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// Logger struct to hold configuration
type Logger struct {
	level  LogLevel
	logger *log.Logger
	file   *os.File
}

// Global logger instance
var logger *Logger

// Initialize the logger (call this once in your program)
func InitLogger(logFile string, level LogLevel) error {
	// Open or create log file
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	// Create logger instance
	logger = &Logger{
		level:  level,
		logger: log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile),
		file:   file,
	}
	return nil
}

// Enhanced version with console output and colors
func Log(level LogLevel, message string, args ...interface{}) {
	if logger == nil {
		log.Printf("[UNINITIALIZED LOGGER] "+message, args...)
		return
	}

	if level >= logger.level {
		var prefix, color string
		switch level {
		case DEBUG:
			prefix = "[DEBUG]"
			color = "\033[36m" // Cyan
		case INFO:
			prefix = "[INFO]"
			color = "\033[32m" // Green
		case WARN:
			prefix = "[WARN]"
			color = "\033[33m" // Yellow
		case ERROR:
			prefix = "[ERROR]"
			color = "\033[31m" // Red
		}

		// File output (no color)
		formattedMsg := fmt.Sprintf(prefix+" "+message, args...)
		logger.logger.Output(2, formattedMsg)

		// Console output (with color)
		reset := "\033[0m"
		fmt.Printf(color+prefix+" "+message+reset+"\n", args...)
	}
}

// Cleanup function to close the log file
func CloseLogger() {
	if logger != nil && logger.file != nil {
		logger.file.Close()
	}
}
