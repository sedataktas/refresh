package runner

import (
	"fmt"
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
	fmt.Println("run with pid : ")
	//fmt.Println(cmd.Stdout)
	//fmt.Println(cmd.Stderr)
	Cmd = cmd
	fmt.Println(cmd.Process.Pid)
}
