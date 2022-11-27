package converter

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

var expectedCommand string = "ffmpeg -i original/0.MTS -c:v libx265 -preset fast -crf 28 -tag:v hvc1 -c:a eac3 -b:a 224k dest/0.mp4 && exiftool -tagsFromFile original/0.MTS -time:all dest/0.mp4"

func TestRightCommand(t *testing.T) {
	err := ConvertVideo("original/0.MTS", "dest/0.mp4", fakeExecCommand)
	// if !<-commandIsCalled {
	// 	t.Error("Command is not called")
	// }
	if err != nil {
		t.Error("Method resulted in error")
	}
}

func getFakeExecCommand(shouldFail bool) ExecCommand {
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

func fakeExecCommand(command string, args...string) *exec.Cmd {
    cs := []string{"-test.run=TestHelperProcess", "--", command}
    cs = append(cs, args...)
    cmd := exec.Command(os.Args[0], cs...)
    cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
    return cmd
}

func TestHelperProcess(t *testing.T){
	fmt.Fprintf(os.Stdout, os.Getenv("GO_WANT_HELPER_PROCESS"))
    if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
        return
    }
	fmt.Fprintf(os.Stdout, "hey")
    os.Exit(0)
}

/*
func TestHelperProcess(t *testing.T) {
	fmt.Println(os.Getenv("GO_WANT_HELPER_PROCESS"))
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	fmt.Println("here")

	if os.Args[3] != "zsh" {
		t.Errorf("Expected command zsh, received %s", os.Args[3])
	}
	if os.Args[4] != "-c" {
		t.Errorf("Expected flag -c, received %s", os.Args[4])
	}
	if os.Args[5] != expectedCommand {
		t.Errorf("Expected command %s but received %s", expectedCommand, os.Args[4])
	}
	
	if os.Getenv("GO_HELPER_PROCESS_SHOULD_FAIL") == "0" {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}
*/
