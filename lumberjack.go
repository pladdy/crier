// Package lumberjack is an opinionaed logger.  It's a wrapper to "log" that provides
// some helpers to make logging a little easier.
//
// It logs with the following log package options:
// log.Ldate, log.Ltime, log.Llongfile
package lumberjack

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	info  *log.Logger
	debug *log.Logger
	error *log.Logger
	warn  *log.Logger
)

// log package uses a callDepth of 2; seems to be what you need to log from
// where log message is calling from
const callDepth = 2

// Debug will only log if the DEBUG environment variable is set.  It logs with a DEBUG prefix.
func Debug(logStatement string, a ...interface{}) {
	if os.ExpandEnv("${DEBUG}") != "" {
		if a != nil {
			debug.Output(callDepth, fmt.Sprintf(logStatement, a...))
		} else {
			debug.Output(callDepth, logStatement)
		}
	}
}

// Error logs statements with an ERROR: prefix
func Error(logStatement string, a ...interface{}) {
	if a != nil {
		error.Output(callDepth, fmt.Sprintf(logStatement, a...))
	} else {
		error.Output(callDepth, logStatement)
	}
}

// Fatal logs with the ERROR prefix and then exits with a non 0 exit status
func Fatal(logStatement string, a ...interface{}) {
	if a != nil {
		error.Output(callDepth, fmt.Sprintf(logStatement, a...))
		os.Exit(1)
	} else {
		error.Output(callDepth, logStatement)
		os.Exit(1)
	}
}

// Hush quiets the logging, point them all to /dev/null
func Hush() {
	devNull, err := os.Create(os.DevNull)
	if err != nil {
		panic("Failed to hush lumberjack!")
	}

	StartLogging(devNull)
}

// Info logs with the INFO: prefix
func Info(logStatement string, a ...interface{}) {
	if a != nil {
		info.Output(callDepth, fmt.Sprintf(logStatement, a...))
	} else {
		info.Output(callDepth, logStatement)
	}
}

// Panic logs statements to error and then calls panic
func Panic(logStatement string, a ...interface{}) {
	if a != nil {
		error.Output(callDepth, fmt.Sprintf(logStatement, a...))
		panic(logStatement)
	} else {
		error.Output(callDepth, logStatement)
		panic(logStatement)
	}
}

// StartLogging starts up the log.Loggers
func StartLogging(out ...io.Writer) {
	if len(out) == 0 {
		info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Llongfile)
		error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
		debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
	} else {
		info = log.New(out[0], "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		warn = log.New(out[0], "WARN: ", log.Ldate|log.Ltime|log.Llongfile)
		error = log.New(out[0], "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
		debug = log.New(out[0], "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
	}
}

// Warn logs statements with the WARN: prefix
func Warn(logStatement string, a ...interface{}) {
	if a != nil {
		warn.Output(callDepth, fmt.Sprintf(logStatement, a...))
	} else {
		warn.Output(callDepth, logStatement)
	}
}
