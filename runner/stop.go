package runner

import (
	"github.com/fatih/color"
)

func Stop() error {
	if err := Cmd.Process.Kill(); err != nil {
		return err
	}

	color.Red("%s --> stopped with:%d\n", GetTime(), Cmd.Process.Pid)
	return nil
}
