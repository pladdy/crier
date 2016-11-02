package lumberjack

import (
	"bytes"
	"os"
	"regexp"
	"strings"
	"testing"
)

// Helpers

// Given a string remove trailing newline
func chomp(s string) string {
	return strings.TrimRight(s, "\n")
}

// Restart logging and return a test buffer
func restartLogging() *bytes.Buffer {
	var testBuffer bytes.Buffer
	StartLogging(&testBuffer)
	return &testBuffer
}

// Given a log string, remove the prefix and newline
func trimLog(level string, s string) string {
	// remove log prefix from buffer
	logPrefix := regexp.MustCompile(level + ":.+?:\\d+: ")
	return chomp(logPrefix.ReplaceAllString(s, ""))
}

// Tests

func TestHush(t *testing.T) {
	// test one call to Hush
	testBuffer := restartLogging()
	Hush()
	expected := ""
	Info("I'm hushed, I can't sing!")

	result := trimLog("INFO", testBuffer.String())
	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	// Test logging, then doing a Hush, then logging again
	testBuffer = restartLogging()
	expected = "Not hushed..."
	Info("Not hushed...")

	result = trimLog("INFO", testBuffer.String())
	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	// Hush again, nothing will get logged
	Hush()
	Info("I got hushed again!")

	// Start logging again, should append
	StartLogging(testBuffer)
	expected += "\nor am I?"
	Info("or am I?")

	result = trimLog("INFO", testBuffer.String())
	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

func TestInfo(t *testing.T) {
	testBuffer := restartLogging()
	expected := "this is info"
	Info(expected)
	result := trimLog("INFO", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	testBuffer = restartLogging()
	expected = "the answer is 42"
	Info("the answer is %v", 42)
	result = trimLog("INFO", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

func TestLumberjack(t *testing.T) {
	StartLogging()

	t.Run("TestWarn", func(t *testing.T) {
		Warn("this is warn")
		Warn("this is a %%v placeholder: %v", 42)
	})

	t.Run("TestError", func(t *testing.T) {
		Error("this is error")
		Error("this is a %%v placeholder: %v", 42)
	})

	t.Run("TestDebug", func(t *testing.T) {
		if os.ExpandEnv("${DEBUG}") != "" {
			os.Unsetenv("${DEBUG}")
		}

		Debug("this is debug") // won't print

		os.Setenv("DEBUG", "1")
		Debug("this is debug") // will print
		Debug("this is a %%v placeholder: %v", 42)
	})

	t.Run("TestPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				Warn("Panic recovered!")
			}
		}()

		Panic("this is panic!")
	})

	t.Run("TestPanicWithVerb", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				Warn("this is a %%v placeholder: %v", 42)
			}
		}()

		Panic("this is panic! %%v placeholder: %v", 42)
	})
}
