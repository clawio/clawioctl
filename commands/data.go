package commands

import (
	"fmt"
	"io"
	"os"

	"github.com/clawio/cli/log"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
)

var DataCommands = cli.Command{
	Name:    "data",
	Aliases: []string{"d"},
	Usage:   "Data commands",
	Subcommands: []cli.Command{
		DownloadCommand,
		UploadCommand,
	},
}
var DownloadCommand = cli.Command{
	Name:      "download",
	Aliases:   []string{"d"},
	Usage:     "Download a BLOB",
	ArgsUsage: "Usage: download <pathspec>",
	Action:    download,
}

func download(c *cli.Context) {
	if c.Args().First() == "" {
		fmt.Println(c.Command.ArgsUsage)
		os.Exit(1)
	}
	sdk := getSDK()
	reader, resp, err := sdk.Data.Download(c.Args().First())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	io.Copy(os.Stdout, reader)
}

var UploadCommand = cli.Command{
	Name:      "upload",
	Aliases:   []string{"u"},
	Usage:     "Upload a BLOB",
	ArgsUsage: "Usage: upload <file> <pathspec>",
	Action:    upload,
}

func upload(c *cli.Context) {
	if len(c.Args()) < 2 {
		fmt.Println(c.Command.ArgsUsage)
		os.Exit(1)
	}
	localFile := c.Args().First()
	pathSpec := c.Args().Get(1)
	sdk := getSDK()
	reader, err := os.Open(localFile)
	if err != nil {
		log.Fatalln(err)
	}
	resp, err := sdk.Data.Upload(pathSpec, reader, "")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	fmt.Println(color.GreenString("BLOB %q uploaded to %q", localFile, pathSpec))
}
