package log

import (
	"log"
	"os"
	"os/user"
	"path"

	"github.com/clawio/cli/config"
	jww "github.com/spf13/jWalterWeatherman"
)

func init() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	home := u.HomeDir
	configDir := path.Join(home, ".clawio")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatalln(err)
	}
	jww.SetLogThreshold(jww.LevelInfo)
	jww.SetLogFile(config.CLILogFile)
	/*
		fd, err := os.OpenFile(config.CLILogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		log.SetOutput(fd)
	*/
}

func Println(v ...interface{}) { jww.INFO.Println(v) }
func Fatalln(v ...interface{}) {
	jww.ERROR.Fatal(v)
}
