package runner

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
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

func WatchCron() {
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

func Watch() {
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Name != buildFileName {
					color.Green("%s --> EVENT! %s\n",
						GetTime(), event.String())

					err := Stop()
					if err != nil {
						log.Fatal(err)
					}
					Build()
					Run()
				}

			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

func addFilesToWatcher() {
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() != buildFileName {
				watcher.Add(path)
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
