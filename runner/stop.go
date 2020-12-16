package runner

import "fmt"

func Stop() error {
	if err := Cmd.Process.Kill(); err != nil {
		return err
	}
	fmt.Printf("stopped with:%d\n", Cmd.Process.Pid)
	return nil
}
