[![Go Report Card](https://goreportcard.com/badge/github.com/pladdy/lumberjack)](https://goreportcard.com/report/github.com/pladdy/lumberjack)

## lumberjack
A Go package for wrapping the "log" package with some helper functions.
Basically it's an opinionated logger (see what I did there?); lumberjacks make
logs so it seemed like a good name.

## Caveat
I'm writing this as yet another Go package that is probably done better
elsewhere.  I'm doing it for practice and my own edification.

## Install
`go get github.com/pladdy/lumberjack`

## Test
`go test`

## Docs
`godoc github.com/pladdy/lumberjack`

## Example
```go
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
```
