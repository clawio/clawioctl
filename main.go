package main

import (
	"fmt"
	"os"

	"github.com/clawio/clawioctl/commands"
	"github.com/codegangsta/cli"
)

// Build information obtained with the help of -ldflags
var (
	buildDate     string // date -u
	gitTag        string // git describe --exact-match HEAD
	gitNearestTag string // git describe --abbrev=0 --tags HEAD
	gitCommit     string // git rev-parse HEAD
)

func main() {

	app := cli.NewApp()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Hugo González Labrador",
			Email: "contact@hugo.labkode.com",
		},
	}
	app.Copyright = "GNU Affero General Public License v3.0"
	app.Name = "clawioctl"
	app.Usage = `
	
	The ClawIO Command Line Interface is the unified tool to manage your ClawIO services.
	`
	app.Commands = []cli.Command{
		commands.ConfigureCommand,
		commands.CleanCommand,
		commands.DataCommands,
		commands.MetaDataCommands,
	}

	app.Version = getVersion(app)

	app.Run(os.Args)
}

func getVersion(app *cli.App) string {
	// if gitTag is not empty we are on release build
	if gitTag != "" {
		return fmt.Sprintf("%s %s commit:%s release-build\n", app.Name, gitNearestTag, gitCommit)
	}
	return fmt.Sprintf("%s %s commit:%s dev-build\n", app.Name, gitNearestTag, gitCommit)
}
