// Lumberjack is an opinionaed logger.  It's a wrapper to "log" that provides
// some helpers to make logging a little easier.
package lumberjack

import (
	"log"
	"os"
)

var (
	info  *log.Logger
	warn  *log.Logger
	error *log.Logger
	debug *log.Logger
)

// Start up the logging handlers; takes option names of new handlers
// Default Handlers:
//    Info, Warn, Error, Debug
func StartLogging() {
	info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Llongfile)
	warn = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Llongfile)
	error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
	debug = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Llongfile)
	Info("Loggers started")
}

func Info(logStatement string) {
	info.Println(logStatement)
}

func Warn(logStatement string) {
	warn.Println(logStatement)
}

func Error(logStatement string) {
	error.Println(logStatement)
}

func Debug(logStatement string) {
	if os.ExpandEnv("${DEBUG}") != "" {
		debug.Println(logStatement)
	}
}

func Panic(logStatement string) {
	error.Panicln(logStatement)
}

func Fatal(logStatement string) {
	error.Fatalln(logStatement)
}
