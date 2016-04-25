package commands

import (
	"fmt"
	"io/ioutil"

	"github.com/clawio/cli/config"
	"github.com/clawio/cli/log"
	"github.com/codegangsta/cli"
	"github.com/ryanuber/columnize"
)

var WhoAmICommand = cli.Command{
	Name:      "whoami",
	Usage:     "Display account",
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
	lines := []string{
		"USERNAME|EMAIL|DISPLAYNAME",
		fmt.Sprintf("%s|%s|%s", user.GetUsername(), user.GetEmail(), user.GetDisplayName()),
	}
	fmt.Println(columnize.SimpleFormat(lines))
}

func getToken() string {
	data, _ := ioutil.ReadFile(config.CLICredentialsFile)
	return string(data)
}
