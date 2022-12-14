package converter

import (
	"fmt"
	"path"
)

type Command interface {
	CombinedOutput() ([]byte, error)
}
type ExecCommand func(command string, args ...string) Command

func ConvertVideo(orig string, dest string, execCommand ExecCommand) error {
	cmd := execCommand(
		"zsh",
		"-c",
		"mkdir -p \""+path.Dir(
			dest,
		)+"\" && ffmpeg -i \""+orig+"\" -c:v libx265 -preset fast -crf 28 -tag:v hvc1 -c:a eac3 -b:a 224k \""+dest+"\" && exiftool -tagsFromFile \""+orig+"\" -time:all \""+dest+"\"",
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
	}
	return err
}
