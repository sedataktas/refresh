package runner

import (
	"errors"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Build() {
	cmd := exec.Command("go", "build", "-o", buildFileName, ".")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		log.Fatal(errors.New(string(errBuf)))
	}
	color.Blue("%s --> app built", GetTime())
}
