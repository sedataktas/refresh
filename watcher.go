package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

var watcher *fsnotify.Watcher

func init() {
	var err error
	watcher, err = fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
}

func watchCron() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			addFilesToWatcher()
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func watch() {
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-watcher.Events:
				//fmt.Printf("EVENT! %#v\n", event)
				cmd := exec.Command("go", "run", "/Users/sedat/go/src/deneme")
				stdout, err := cmd.Output()

				if err != nil {
					fmt.Println(err.Error())
				}

				fmt.Println(string(stdout))
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func addFilesToWatcher() {
	err := filepath.Walk("/Users/sedat/go/src/deneme",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			watcher.Add(path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
