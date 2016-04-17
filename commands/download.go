package commands

import (
	"github.com/clawio/cli/log"
	"github.com/codegangsta/cli"
	"io"
	"os"
)

var DownloadCommand = cli.Command{
	Name:      "download",
	Aliases:   []string{"d"},
	Usage:     "Download a BLOBL",
	ArgsUsage: "",
	Action:    download,
}

func download(c *cli.Context) {
	sdk := getSDK()
	reader, resp, err := sdk.Data.Download("myblob")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	io.Copy(os.Stdout, reader)
}
