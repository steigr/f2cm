package watch

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
)

var (
	watchers = map[string]*fsnotify.Watcher{}
)

func Init(directory string) (events chan fsnotify.Event, errors chan error, err error) {
	watchers[directory], err = fsnotify.NewWatcher()
	log.Println("Watching", directory)
	filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		watchers[directory].Add(path)
		return nil
	})
	return watchers[directory].Events, watchers[directory].Errors, err
}

func Close() {
	for directory, watcher := range watchers {
		log.Println("Stop watching", directory)
		watcher.Close()
	}
}
