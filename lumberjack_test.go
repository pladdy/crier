package lumberjack

import (
	"os"
	"testing"
)

func TestLumberjack(t *testing.T) {
	StartLogging()

	t.Run("TestInfo", func(t *testing.T) { Info("this is info") })
	t.Run("TestWarn", func(t *testing.T) { Warn("this is warn") })
	t.Run("TestError", func(t *testing.T) { Error("this is error") })
	t.Run("TestDebug", func(t *testing.T) {
		if os.ExpandEnv("${DEBUG}") != "" {
			os.Unsetenv("${DEBUG}")
		}

		Debug("this is debug") // won't print

		os.Setenv("DEBUG", "1")
		Debug("this is debug") // will print
	})
	t.Run("TestPanic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				Warn("Panic recovered!")
			}
		}()

		Panic("this is panic!")
	})
}
