package main

import (
	"github.com/117/logger"
)

// Destroy let's you kill daemons with regex matching.
func Destroy(args ...string) {
	if len(args) < 1 {
		logger.Log(logger.Err, "No expression or PID was provided.")
		return
	}

	var (
		expression string
		killed     int
	)

	if len(args) > 0 {
		expression = args[0]
	}

	for _, demon := range FindMatchingDemons(expression) {
		if MatchDemon(demon, expression) {
			demon.Kill()
			killed++
		}
	}

	logger.Logf(logger.Info, "Found and destroyed %d demon(s).", killed)
}
