package args

import "github.com/spf13/pflag"

var (
	Debug  = pflag.BoolP("debug", "D", false, "Enable debug messages")
	Source = pflag.BoolP("source", "s", false, "Adds the source into log messages")
)

func init() {
	pflag.Parse()
}
