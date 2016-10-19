package lumberjack

import (
	"os"
	"testing"
)

func TestLumberjack(t *testing.T) {
	StartLogging()

	t.Run("TestInfo", func(t *testing.T) {
		Info("this is info")
		Info("this is a %%v placeholder: %v", 42)
		Info("this is a number: %v and this is a string: %s", 42, "banana")
	})

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
