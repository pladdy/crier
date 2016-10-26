// Lumberjack is an opinionaed logger.  It's a wrapper to "log" that provides
// some helpers to make logging a little easier.
package lumberjack

import (
	"fmt"
	"log"
	"os"
)

var (
	info           *log.Logger
	warn           *log.Logger
	error          *log.Logger
	debug          *log.Logger
	loggersStarted bool
	hushed         bool
)

// log package uses a callDepth of 2; if we don't bump to 3 when this
// package gets used the file shown will always be this one
const callDepth = 2

// Start up the logging handlers; only initializes once
// Default Handlers:
//    Info, Warn, Error, Debug
func StartLogging() {
	// loggers aren't already started, or the logging is hushed, start
	if loggersStarted != true || hushed != false {
		info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Llongfile)
		error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
		debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
		loggersStarted = true
		hushed = false
		Info("Loggers started")
	}
}

// Quiet the logging, point them all to /dev/null
func Hush() {
	devNull, err := os.Create(os.DevNull)
	if err != nil {
		panic("Failed to hush lumberjack!")
	}

	info = log.New(devNull, "", log.Ldate)
	warn = log.New(devNull, "", log.Ldate)
	error = log.New(devNull, "", log.Ldate)
	debug = log.New(devNull, "", log.Ldate)
	loggersStarted = true
	hushed = true
}

func Info(logStatement string, a ...interface{}) {
	if a != nil {
		info.Output(callDepth, fmt.Sprintf(logStatement, a...))
	} else {
		info.Output(callDepth, logStatement)
	}
}

func Warn(logStatement string, a ...interface{}) {
	if a != nil {
		warn.Output(callDepth, fmt.Sprintf(logStatement, a...))
	} else {
		warn.Output(callDepth, logStatement)
	}
}

func Error(logStatement string, a ...interface{}) {
	if a != nil {
		error.Output(callDepth, fmt.Sprintf(logStatement, a...))
	} else {
		error.Output(callDepth, logStatement)
	}
}

func Debug(logStatement string, a ...interface{}) {
	if os.ExpandEnv("${DEBUG}") != "" {
		if a != nil {
			debug.Output(callDepth, fmt.Sprintf(logStatement, a...))
		} else {
			debug.Output(callDepth, logStatement)
		}
	}
}

func Panic(logStatement string, a ...interface{}) {
	if a != nil {
		error.Output(callDepth, fmt.Sprintf(logStatement, a...))
		panic(logStatement)
	} else {
		error.Output(callDepth, logStatement)
		panic(logStatement)
	}
}

func Fatal(logStatement string, a ...interface{}) {
	if a != nil {
		error.Output(callDepth, fmt.Sprintf(logStatement, a...))
		os.Exit(1)
	} else {
		error.Output(callDepth, logStatement)
		os.Exit(1)
	}
}
