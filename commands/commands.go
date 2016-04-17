package commands

import (
	"github.com/clawio/cli/client"
	"github.com/clawio/cli/config"
	"github.com/clawio/sdk"
)

func getSDK() *sdk.SDK {
	cfg := config.Get()
	clientConfig := &client.Config{ClientID: cfg.Username, ClientSecret: cfg.Password}
	tokenStore := client.NewFileTokenStore(config.CLICredentialsFile)
	c := client.NewClient(clientConfig, tokenStore)
	urls := &sdk.ServiceEndpoints{
		AuthServiceBaseURL: cfg.AuthenticationServiceBaseURL,
		DataServiceBaseURL: cfg.DataServiceBaseURL,
	}
	s := sdk.New(urls, c)
	return s
}
