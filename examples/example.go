package main

import (
	"fmt"

	"github.com/pladdy/lumberjack"
)

func main() {
	fmt.Println("Firing up the lumberjack!")

	lumberjack.StartLogging()
	lumberjack.Info("I'm a lumberjack and I'm okay")
	lumberjack.Warn("I eat all night and I work all day")
	lumberjack.Debug("I cut down trees, I eat my lunch")
	lumberjack.Error("I go to the lavotry")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovering Panic")
		}
	}()

	lumberjack.Panic("Timber!!!")
}
