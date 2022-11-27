package converter_test

import (
	"os"
	"os/exec"
	"testing"
)

var expectedCommand string = "ffmpeg -i original/0.MTS -c:v libx265 -preset fast -crf 28 -tag:v hvc1 -c:a eac3 -b:a 224k dest/0.mp4 && exiftool -tagsFromFile original/0.MTS -time:all dest/0.mp4"

func TestRightCommand(t *testing.T) {
	execCommand := getFakeExecCommand(false)
	err := ConvertVideo("original/0.MTS", "dest/0.mp4", execCommand)

}

type fakeExecCommand func(command string, args...string) *exec.Cmd

func getFakeExecCommand(shouldFail bool) fakeExecCommand {
	val := "0"
	if shouldFail {
		val = "1"
	}
	return func(command string, args...string) *exec.Cmd {
		cs := []string{"-test.run=TestHelperProcess", "--", command}
		cs = append(cs, args...)
		cmd := exec.Command(os.Args[0], cs...)
		cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", "GO_HELPER_PROCESS_SHOULD_FAIL=" + val}
		return cmd
	}
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}

	if os.Args[
	
	if os.Getenv("GO_HELPER_PROCESS_SHOULD_FAIL") == "0" {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
