package log

import (
	"log"
	"os"
	"os/user"
	"path"

	"github.com/clawio/clawioctl/config"
	"github.com/fatih/color"
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
}

func Println(v ...interface{}) { jww.INFO.Println(v) }
func Fatalln(v ...interface{}) {
	jww.ERROR.Fatal(color.RedString("%+v", v))
}
