package converter

type Command interface {
	Run() error
}
type ExecCommand func(command string, args...string) Command

func ConvertVideo(orig string, dest string, execCommand ExecCommand) error {
	cmd := execCommand("zsh", "-c", "ffmpeg -i " + orig + " -c:v libx265 -preset fast -crf 28 -tag:v hvc1 -c:a eac3 -b:a 224k " + dest + " && exiftool -tagsFromFile original/0.MTS -time:all dest/0.mp4")
	err := cmd.Run()
	return err
}
