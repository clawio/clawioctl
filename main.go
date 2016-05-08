package main

import (
	"os"

	"github.com/clawio/clawioctl/commands"
	"github.com/codegangsta/cli"
)

var VERSION string

func main() {

	app := cli.NewApp()
	app.Version = VERSION
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Hugo González Labrador",
			Email: "contact@hugo.labkode.com",
		},
	}
	app.Copyright = "GNU Affero General Public License v3.0"
	app.Name = "clawio"
	app.Usage = `
	
	The ClawIO Command Line Interface is the unified tool to manage your ClawIO services.
	`
	app.Commands = []cli.Command{
		commands.ConfigureCommand,
		commands.CleanCommand,
		commands.DataCommands,
		commands.MetaDataCommands,
	}

	app.Run(os.Args)
}
