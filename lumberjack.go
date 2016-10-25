// Lumberjack is an opinionaed logger.  It's a wrapper to "log" that provides
// some helpers to make logging a little easier.
package lumberjack

import (
	"log"
	"os"
)

var (
	info           *log.Logger
	warn           *log.Logger
	error          *log.Logger
	debug          *log.Logger
	loggersStarted bool
)

// Start up the logging handlers; only initializes once
// Default Handlers:
//    Info, Warn, Error, Debug
func StartLogging() {
	if loggersStarted != true {
		info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
		warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Llongfile)
		error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
		debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
		loggersStarted = true
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
	loggersStarted = false
}

func Info(logStatement string, a ...interface{}) {
	if a != nil {
		info.Printf(logStatement+"\n", a...)
	} else {
		info.Println(logStatement)
	}
}

func Warn(logStatement string, a ...interface{}) {
	if a != nil {
		warn.Printf(logStatement+"\n", a...)
	} else {
		warn.Println(logStatement)
	}
}

func Error(logStatement string, a ...interface{}) {
	if a != nil {
		error.Printf(logStatement+"\n", a...)
	} else {
		error.Println(logStatement)
	}
}

func Debug(logStatement string, a ...interface{}) {
	if os.ExpandEnv("${DEBUG}") != "" {
		if a != nil {
			debug.Printf(logStatement+"\n", a...)
		} else {
			debug.Println(logStatement)
		}
	}
}

func Panic(logStatement string, a ...interface{}) {
	if a != nil {
		error.Panicf(logStatement+"\n", a...)
	} else {
		error.Panicln(logStatement)
	}
}

func Fatal(logStatement string, a ...interface{}) {
	if a != nil {
		error.Fatalf(logStatement+"\n", a...)
	} else {
		error.Fatalln(logStatement)
	}
}
