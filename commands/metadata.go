package commands

import (
	"fmt"
	"os"

	"github.com/clawio/clawioctl/log"
	"github.com/codegangsta/cli"
	"github.com/fatih/color"
	"github.com/ryanuber/columnize"
)

var MetaDataCommands = cli.Command{
	Name:    "metadata",
	Aliases: []string{"m"},
	Usage:   "Metadata commands",
	Subcommands: []cli.Command{
		ExamineObjectCommand,
		ListTreeCommand,
		InitCommand,
		DeleteObjectCommand,
		MoveObjectCommand,
	},
}

var InitCommand = cli.Command{
	Name:      "init",
	Aliases:   []string{"i"},
	Usage:     "Init home tree",
	ArgsUsage: "Usage: init",
	Action:    init2,
}

var ExamineObjectCommand = cli.Command{
	Name:      "examine",
	Aliases:   []string{"e"},
	Usage:     "Examine an object",
	ArgsUsage: "Usage: examine <pathspec>",
	Action:    examineObject,
}

var ListTreeCommand = cli.Command{
	Name:      "ls",
	Aliases:   []string{"l"},
	Usage:     "List a tree",
	ArgsUsage: "Usage: ls <pathspec>",
	Action:    listTree,
}
var DeleteObjectCommand = cli.Command{
	Name:      "rm",
	Aliases:   []string{"r"},
	Usage:     "Delete an object",
	ArgsUsage: "Usage: rm <pathspec>",
	Action:    deleteObject,
}

var MoveObjectCommand = cli.Command{
	Name:      "mv",
	Aliases:   []string{"r"},
	Usage:     "Move an object",
	ArgsUsage: "Usage: mv <sourcepathspec> <targetpathspec>",
	Action:    moveObject,
}

func examineObject(c *cli.Context) {
	if c.Args().First() == "" {
		fmt.Println(c.Command.ArgsUsage)
		os.Exit(1)
	}
	sdk := getSDK()
	info, resp, err := sdk.Meta.ExamineObject(c.Args().First())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	lines := []string{"PATHSPEC|TYPE|SIZE|MIMETYPE|CHECKSUM"}
	line := fmt.Sprintf("%s|%d|%d|%s|%s",
		info.GetPathSpec(), info.GetType(), info.GetSize(), info.GetMimeType(), info.GetChecksum())
	lines = append(lines, line)
	fmt.Println(columnize.SimpleFormat(lines))
}

func listTree(c *cli.Context) {
	if c.Args().First() == "" {
		fmt.Println(c.Command.ArgsUsage)
		os.Exit(1)
	}
	sdk := getSDK()
	infos, resp, err := sdk.Meta.ListTree(c.Args().First())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	lines := []string{"PATHSPEC|TYPE|SIZE|MIMETYPE|CHECKSUM"}
	for _, info := range infos {
		line := fmt.Sprintf("%s|%d|%d|%s|%s",
			info.GetPathSpec(), info.GetType(), info.GetSize(), info.GetMimeType(), info.GetChecksum())
		lines = append(lines, line)
	}
	fmt.Println(columnize.SimpleFormat(lines))
}

func init2(c *cli.Context) {
	sdk := getSDK()
	resp, err := sdk.Meta.Init()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	fmt.Println(color.GreenString("%s", "Home tree created!"))
}

func deleteObject(c *cli.Context) {
	if c.Args().First() == "" {
		fmt.Println(c.Command.ArgsUsage)
		os.Exit(1)
	}
	sdk := getSDK()
	resp, err := sdk.Meta.DeleteObject(c.Args().First())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	fmt.Println(color.GreenString("Deleted object %q!", c.Args().First()))
}

func moveObject(c *cli.Context) {
	if len(c.Args()) < 2 {
		fmt.Println(c.Command.ArgsUsage)
		os.Exit(1)
	}
	sourcePath := c.Args().First()
	targetPath := c.Args().Get(1)
	sdk := getSDK()
	resp, err := sdk.Meta.MoveObject(sourcePath, targetPath)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(resp)
	fmt.Println(color.GreenString("Moved object from %q to %q!", sourcePath, targetPath))
}
