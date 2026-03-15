package logger

import (
	"github.com/charmbracelet/log"
	"os"

	"github.com/CyberDruga/twitter-bot/src/args"
)

func init() {

	level := log.InfoLevel

	if *args.Debug {
		level = log.DebugLevel
	}

	log.SetLevel(level)
	log.SetReportCaller(*args.Source)
	log.SetOutput(os.Stderr)

	// NOTE: detects if it's being run on systemd. No point in showing timestamp twice.
	log.SetReportTimestamp(os.Getenv("INVOCATION_ID") == "")

}
