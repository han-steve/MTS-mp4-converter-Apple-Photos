package converter

import "os/exec"

type ExecCommand func(command string, args...string) *exec.Cmd

func ConvertVideo(orig string, dest string, execCommand ExecCommand) error {
	cmd := execCommand("zsh", "-c")
	cmd.CombinedOutput()
	return nil
}
