package watch

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/fsnotify/fsnotify"
	"github.com/jesseobrien/torque-cli/internal/run"
	"github.com/spf13/cobra"
)

var (
	WatchCmd = &cobra.Command{
		Use:   "watch",
		Short: "Watch and re-run a Torque app. Use --build to output a binary as well.",
		Long:  "",
		Run:   executeWatch,
	}
)

func executeWatch(cmd *cobra.Command, args []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalf("NewWatcher failed: %s ", err.Error())
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		defer close(done)

		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Infof("%s %s\n", event.Name, event.Op)

				// @TODO: Figure out how to run this in a channel blocked thread or something
				// so we can kill it when a new event comes in
				if err := run.RunCmd.Execute(); err != nil {
					log.Errorf("error running 'run' %s", err.Error())
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Errorf("error:", err)
			}
		}

	}()

	if err := filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && !strings.Contains(path, ".git") && !strings.Contains(path, "dist") {
			log.Infof("adding path to watcher: %s", path)
			return watcher.Add(path)
		}

		return nil
	}); err != nil {
		log.Fatalf("Filepath walk failed: %s", err.Error())
		return
	}

	err = watcher.Add("./")
	if err != nil {
		log.Fatalf("Add failed: %s", err.Error())
	}
	<-done

}
