package commands

import (
	"os"

	"github.com/clawio/cli/config"
	"github.com/clawio/cli/log"
	"github.com/codegangsta/cli"
)

var CleanCommand = cli.Command{
	Name:        "clean",
	Usage:       "Clean configuration",
	Description: "Removes ClawIO configuration files (~/.clawio/*) to start with a fresh installation",
	ArgsUsage:   "",
	Action:      clean,
}

func clean(c *cli.Context) {
	if err := os.RemoveAll(config.CLIConfigDir); err != nil {
		log.Fatalln(err)
	}
}
