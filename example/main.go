package main

import (
	"fmt"
	log "github.com/turgay/logger"
)

func main() {
	logger := log.Logger{}

	logger.SetLevelWithDefault("trace", "info")
	printAllLevels(logger, "Prints 6 lines")

	logger.SetLevelWithDefault("info", "info")
	printAllLevels(logger, "Prints 5 lines")

	logger.SetLevelWithDefault("debug", "info")
	printAllLevels(logger, "Prints 4 lines")

	logger.SetLevelWithDefault("warn", "info")
	printAllLevels(logger, "Prints 3 lines")

	logger.SetLevelWithDefault("error", "info")
	printAllLevels(logger, "Prints 2 lines")

	logger.SetLevelWithDefault("fatal", "info")
	printAllLevels(logger, "Prints 1 line")

}

func printAllLevels(logger log.Logger, msg string) {
	logger.Info(msg)
	logger.Debug(msg)
	logger.Trace(msg)
	logger.Warn(msg)
	logger.Error(msg)
	logger.Fatal(msg)
	fmt.Println()
}
