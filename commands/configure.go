package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/clawio/clawioctl/config"
	"github.com/clawio/clawioctl/log"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

var ConfigureCommand = cli.Command{
	Name:  "configure",
	Usage: "Configure ClawIO CLI options",
	Description: `
 Configure ClawIO  CLI options. If  this command is  run with no  arguments,
 you will be  prompted for configuration values such as  your ClawIO credentials
 (username  and password).  If  your config  file does  not  exist (the  default
 location is ~/.clawio/config),  the ClawIO CLI will create it  for you. To keep
 an existing value, hit enter when prompted for the value. When you are propmted
 for information,  the current  value will  be displayed  in [brackets].  If the
 config  item has  no value,  it  will be  displayed  as [None].  Note that  the
 configure command only work  with values from the config file.  It does not use
 any values from environment variables.

 Note: the ClawIO Access Token  obtained after validating the ClawIO Credentials
 will be written to the shared credentials file (~/.clawio/credentials).

 Note: ClawIO will log additional information to a log file (default location is
 ~/.clawio/log).
`,
	ArgsUsage: "",
	Action:    configure,
}

func configure(c *cli.Context) {
	cfg := ask()
	config.Set(cfg)

	sdk := getSDK()
	token, resp, err := sdk.Auth.Authenticate(cfg.Username, cfg.Password)
	log.Println(resp)
	config.SetToken(token)
	if err != nil {
		log.Fatalln(err)
	}
	config.Set(cfg)
	fmt.Println(color.GreenString("Configuration saved to %q", config.CLIConfigFile))
}

func ask() *config.Config {
	cfg := config.Get()
	if cfg.AuthenticationServiceBaseURL == "" {
		cfg.AuthenticationServiceBaseURL = config.DefaultAuthenticationServiceBaseURL
	}
	if cfg.DataServiceBaseURL == "" {
		cfg.DataServiceBaseURL = config.DefaultDataServiceBaseURL
	}
	if cfg.MetaDataServiceBaseURL == "" {
		cfg.MetaDataServiceBaseURL = config.DefaultMetaDataServiceBaseURL
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter username (%s): ", cfg.Username)
	username, _ := reader.ReadString('\n')

	fmt.Printf("Enter password (******): ")
	bytePassword, _ := terminal.ReadPassword(int(syscall.Stdin))
	password := string(bytePassword)
	fmt.Println()

	fmt.Printf("Enter authentication service base URL (%s): ", cfg.AuthenticationServiceBaseURL)
	authenticationServiceBaseURL, _ := reader.ReadString('\n')

	fmt.Printf("Enter data service base URL (%s): ", cfg.DataServiceBaseURL)
	dataServiceBaseURL, _ := reader.ReadString('\n')

	fmt.Printf("Enter metadata service base URL (%s): ", cfg.MetaDataServiceBaseURL)
	metaDataServiceBaseURL, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)
	authenticationServiceBaseURL = strings.TrimSpace(authenticationServiceBaseURL)
	dataServiceBaseURL = strings.TrimSpace(dataServiceBaseURL)
	metaDataServiceBaseURL = strings.TrimSpace(metaDataServiceBaseURL)

	if username != "" {
		cfg.Username = username
	}
	if password != "" {
		cfg.Password = password
	}
	if authenticationServiceBaseURL != "" {
		cfg.AuthenticationServiceBaseURL = authenticationServiceBaseURL
	}
	if dataServiceBaseURL != "" {
		cfg.DataServiceBaseURL = dataServiceBaseURL
	}
	if metaDataServiceBaseURL != "" {
		cfg.MetaDataServiceBaseURL = metaDataServiceBaseURL
	}
	return cfg
}
