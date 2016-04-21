package commands

import (
	"github.com/clawio/cli/client"
	"github.com/clawio/cli/config"
	"github.com/clawio/sdk"
)

func getSDK() *sdk.SDK {
	cfg := config.Get()
	clientCredentials := &client.Credentials{ClientID: cfg.Username, ClientSecret: cfg.Password}
	tokenStore := client.NewFileTokenStore(config.CLICredentialsFile)
	c := client.NewClientWithAuth(clientCredentials, tokenStore)
	urls := &sdk.ServiceEndpoints{
		AuthServiceBaseURL: cfg.AuthenticationServiceBaseURL,
		DataServiceBaseURL: cfg.DataServiceBaseURL,
	}
	s := sdk.New(urls, c)
	return s
}
