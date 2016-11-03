package lumberjack

import (
	"bytes"
	"fmt"
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
func newLoggingBuffer() *bytes.Buffer {
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

func TestHushOnce(t *testing.T) {
	// test one call to Hush
	testBuffer := newLoggingBuffer()
	Hush()
	expected := ""
	Info("I'm hushed, I can't sing!")

	result := trimLog("INFO", testBuffer.String())
	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

// make sure logging starts again after a hush
func TestHushThenLog(t *testing.T) {
	// Test logging, then doing a Hush, then logging again
	testBuffer := newLoggingBuffer()
	expected := "Not hushed..."
	Info("Not hushed...")

	result := trimLog("INFO", testBuffer.String())
	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	// Hush, nothing will get logged
	Hush()
	Info("I got hushed!")

	// Start logging again, should append
	StartLogging(testBuffer)
	expected += "\nNot hushed anymore?"
	Info("Not hushed anymore?")

	result = trimLog("INFO", testBuffer.String())
	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

func TestDebug(t *testing.T) {
	os.Setenv("DEBUG", "1")

	testBuffer := newLoggingBuffer()
	expected := "this is Debug"
	Debug(expected)
	result := trimLog("DEBUG", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	testBuffer = newLoggingBuffer()
	expected = "the answer is 42"
	Debug("the answer is %v", 42)
	result = trimLog("DEBUG", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	os.Unsetenv("DEBUG")
}

func TestError(t *testing.T) {
	testBuffer := newLoggingBuffer()
	expected := "this is Error"
	Error(expected)
	result := trimLog("ERROR", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	testBuffer = newLoggingBuffer()
	expected = "the answer is 42"
	Error("the answer is %v", 42)
	result = trimLog("ERROR", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

func TestInfo(t *testing.T) {
	testBuffer := newLoggingBuffer()
	expected := "this is info"
	Info(expected)
	result := trimLog("INFO", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	testBuffer = newLoggingBuffer()
	expected = "the answer is 42"
	Info("the answer is %v", 42)
	result = trimLog("INFO", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

func TestWarn(t *testing.T) {
	testBuffer := newLoggingBuffer()
	expected := "this is Warn"
	Warn(expected)
	result := trimLog("WARN", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	testBuffer = newLoggingBuffer()
	expected = "the answer is 42"
	Warn("the answer is %v", 42)
	result = trimLog("WARN", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}

func TestPanic(t *testing.T) {
	testBuffer := newLoggingBuffer()
	expected := "this is Panic"

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected recover to catch a panic")
		}
	}()
	Panic(expected)

	result := trimLog("ERROR", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}

	// test with an interface
	testBuffer = newLoggingBuffer()
	expected = "the answer is 42"

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering: %v", r)
		}
	}()
	Panic("the answer is %v", 42)
	result = trimLog("ERROR", testBuffer.String())

	if result != expected {
		t.Error("Got:", result, "Expected:", expected)
	}
}
