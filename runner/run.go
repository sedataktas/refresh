package runner

import (
	"github.com/fatih/color"
	"log"
	"os"
	"os/exec"
)

var Cmd *exec.Cmd

const (
	buildFileName = "build"
)

func Run() {
	run := "./" + buildFileName
	cmd := exec.Command(run)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		log.Fatal(err.Error())
	}

	color.Cyan("%s --> run with pid : %v", GetTime(), cmd.Process.Pid)
	color.Cyan("--------------------")

	Cmd = cmd
}
