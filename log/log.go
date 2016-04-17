package log

import (
	"fmt"
	"github.com/clawio/cli/config"
	"log"
	"os"
	"os/user"
	"path"
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
	fd, err := os.OpenFile(config.CLILogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(fd)
}

func Println(v ...interface{}) { log.Println(v) }
func Fatalln(v ...interface{}) {
	fmt.Printf("Error performing the operation. See %q for more details.", config.CLILogFile)
	log.Fatal(v)
}
