package commands

import (
	"github.com/117/logger"
	"github.com/streamwithme/demon/helpers"
)

// Destroy let's you kill daemons with regex matching.
func Destroy(args ...string) {
	if len(args) < 2 {
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

	for _, demon := range helpers.FindMatchingDemons(expression) {
		if helpers.MatchDemon(demon, expression) {
			demon.Kill()
			killed++
		}
	}

	logger.Logf(logger.Info, "Found and destroyed %d demon(s).", killed)
}
