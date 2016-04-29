package commands

import (
	"github.com/clawio/clawioctl/client"
	"github.com/clawio/clawioctl/config"
	"github.com/clawio/sdk"
)

func getSDK() *sdk.SDK {
	cfg := config.Get()
	clientCredentials := &client.Credentials{AuthenticationServiceBaseURL: cfg.AuthenticationServiceBaseURL, ClientID: cfg.Username, ClientSecret: cfg.Password}
	tokenStore := client.NewFileTokenStore(config.CLICredentialsFile)
	c := client.NewClientWithAuth(clientCredentials, tokenStore)
	urls := &sdk.ServiceEndpoints{
		AuthServiceBaseURL:     cfg.AuthenticationServiceBaseURL,
		DataServiceBaseURL:     cfg.DataServiceBaseURL,
		MetaDataServiceBaseURL: cfg.MetaDataServiceBaseURL,
	}
	s := sdk.New(urls, c)
	return s
}
