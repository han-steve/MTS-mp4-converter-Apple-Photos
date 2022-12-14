package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/SteveHan-233/MTS-to-mp4/pkg/converter"
	"github.com/SteveHan-233/MTS-to-mp4/pkg/traverser"
	"github.com/schollz/progressbar/v3"
)

func realExecCommand(command string, args ...string) converter.Command {
	return exec.Command(command, args...)
}

func main() {
	if len(os.Args) != 4 {
		fmt.Println(
			`To execute this command, provide the root directory, source folder name, and destination folder name as command line arguments.
			For example, you can provide arguments /Users/user/Desktop /Source /Destination.`,
		)
		os.Exit(1)
	}
	root := os.Args[1]
	list := traverser.GetConversionList(
		os.DirFS(root),
		os.Args[2],
		os.Args[3],
	)
	bar := progressbar.Default(int64(len(list)))
	for _, item := range list {
		bar.Describe("Converting video " + item.Source)
		err := converter.ConvertVideo(
			path.Join(root, item.Source),
			path.Join(root, item.Dest),
			realExecCommand,
		)
		if err != nil {
			fmt.Println(err)
		}
		bar.Add(1)
	}
}
