package runner

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func Build() {
	cmd := exec.Command("go", "build", "-o", buildFileName, ".")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		fmt.Println(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(os.Stdout, stdout)
	ioutil.ReadAll(stderr)
	fmt.Println("built")
	err = cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}
}
