package commands

import (
	"fmt"
	"github.com/clawio/cli/config"
	"github.com/clawio/cli/log"
	"github.com/codegangsta/cli"
	"io/ioutil"
)

var WhoAmICommand = cli.Command{
	Name:      "whoami",
	Aliases:   []string{"who"},
	Usage:     "Display information of current logged in user",
	ArgsUsage: "",
	Action:    whoami,
}

func whoami(c *cli.Context) {
	sdk := getSDK()
	user, resp, err := sdk.Auth.Verify(getToken())

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	fmt.Printf("Username: %s\n", user.GetUsername())
	fmt.Printf("Email: %s\n", user.GetEmail())
	fmt.Printf("Display name: %s\n", user.GetDisplayName())
}

func getToken() string {
	data, _ := ioutil.ReadFile(config.CLICredentialsFile)
	return string(data)
}
