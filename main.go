package main

import (
	"github.com/fatih/color"
	"refresh/runner"
)

func main() {
	runner.Build()
	runner.Run()
	color.White("Delay time 5 seconds")
	quit := make(chan struct{})
	go runner.WatchCron()
	runner.Watch()
	<-quit
}
