//go:build !js
// +build !js

package base

import (
	"fmt"
	"os"

	"biyard.co/common/logger"
	"github.com/fsnotify/fsnotify"
)

func (r *Server) WatchDir(path string) {
	entries, err := os.ReadDir(path)

	for _, d := range entries {
		if d.IsDir() && r.Filter(d.Name()) {
			n := fmt.Sprintf("%s/%s", path, d.Name())

			err = r.watcher.Add(n)
			if err != nil {
				panic(err)
			}
			r.WatchDir(n)
		}
	}
}

func watchFilesImpl(r *Server) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	r.watcher = watcher
	watcher.Add(".")

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					logger.Infof(nil, "Config file changed: %s", event.Name)
					r.Stop()
				}
			case err, ok := <-watcher.Errors:
				logger.Errorf(nil, "Watcher error: %v %v", err, ok)
			}
		}
	}()
	r.WatchDir(".")

	logger.Infof(nil, "enabled to watch file changes")
}

func init() {
	WatchFiles = watchFilesImpl
}
