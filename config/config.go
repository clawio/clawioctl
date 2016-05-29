package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
	"path"
)

const (
	DefaultAuthenticationServiceBaseURL = "http://localhost:1502/api/v1/authentication/"
	DefaultDataServiceBaseURL           = "http://localhost:1502/api/v1/data/"
	DefaultMetaDataServiceBaseURL       = "http://localhost:1502/api/v1/metadata/"
)

var CLIConfigDir string
var CLILogFile string
var CLICredentialsFile string
var CLIConfigFile string

type Config struct {
	Username                     string
	Password                     string
	AuthenticationServiceBaseURL string
	DataServiceBaseURL           string
	MetaDataServiceBaseURL       string
}

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := u.HomeDir
	CLIConfigDir = path.Join(home, ".clawio")
	CLIConfigFile = path.Join(CLIConfigDir, "config")
	CLILogFile = path.Join(CLIConfigDir, "log")
	CLICredentialsFile = path.Join(CLIConfigDir, "credentials")
}

func Get() *Config {
	c := &Config{}
	data, err := ioutil.ReadFile(CLIConfigFile)
	if err != nil {
		return c
	}
	if err := json.Unmarshal(data, c); err != nil {
		return c
	}
	return c
}

func Set(cfg *Config) {
	data, _ := json.MarshalIndent(cfg, "", "  ")
	ioutil.WriteFile(CLIConfigFile, data, 0600)
}

func GetToken() string {
	data, err := ioutil.ReadFile(CLICredentialsFile)
	if err != nil {
		return ""
	}
	return string(data)
}

func SetToken(token string) {
	if err := ioutil.WriteFile(CLICredentialsFile, []byte(token), 0600); err != nil {
		log.Fatalln(err)
	}
}
