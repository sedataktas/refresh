package main

import (
	"refresh/runner"
)

func main() {
	runner.Build()
	runner.Run()
	quit := make(chan struct{})
	go runner.WatchCron()
	runner.Watch()
	<-quit
}
