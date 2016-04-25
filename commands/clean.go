package commands

import (
	"fmt"
	"os"

	"github.com/clawio/cli/config"
	"github.com/clawio/cli/log"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
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
	fmt.Println(color.GreenString("%s", "Configuration cleaned!"))
}
