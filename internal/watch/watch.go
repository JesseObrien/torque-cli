package watch

import (
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/apex/log"
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
	// Ensure modd is installed
	// if not, install it
	path, err := exec.LookPath("modd")
	if err != nil || path == "" {
		log.Info("🔨 modd is not installed, installing...")
		instCmd := exec.CommandContext(cmd.Context(), "go", "install", "github.com/cortesi/modd/cmd/modd@latest")
		instCmd.Stdout = os.Stdout
		instCmd.Stderr = os.Stderr

		home, err := os.UserHomeDir()
		if err != nil {
			log.Fatalf("could not obtain user home directory: %s", err.Error())
			return
		}

		instCmd.Dir = home
		if err := instCmd.Run(); err != nil {
			log.Fatal(err.Error())
			return
		}

		log.Info("🔨 modd is good to go.")
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	moddCmd := exec.CommandContext(cmd.Context(), "modd")
	moddCmd.Stdout = os.Stdout
	moddCmd.Stderr = os.Stderr

	if err := moddCmd.Run(); err != nil {
		log.Fatalf("modd failed to run: %s", err.Error())
		return
	}

	<-done

	// Run modd with the project config

	// watcher, err := fsnotify.NewWatcher()
	// if err != nil {
	// 	log.Fatalf("NewWatcher failed: %s ", err.Error())
	// }
	// defer watcher.Close()

	// done := make(chan bool)
	// go func() {
	// 	defer close(done)

	// 	ctx := context.Background()

	// 	for {
	// 		select {
	// 		case event, ok := <-watcher.Events:
	// 			if !ok {
	// 				return
	// 			}

	// 			log.Infof("%s %s\n", event.Name, event.Op)
	// 			ctx.Done()

	// 			ctx = context.Background()

	// 			// @TODO: Figure out how to run this in a channel blocked thread or something
	// 			// so we can kill it when a new event comes in
	// 			go func(ctxt context.Context) {
	// 				if err := run.RunCmd.ExecuteContext(ctxt); err != nil {
	// 					log.Errorf("error running 'run' %s", err.Error())
	// 				}
	// 			}(ctx)
	// 		case err, ok := <-watcher.Errors:
	// 			if !ok {
	// 				return
	// 			}
	// 			log.Errorf("error:", err)
	// 		}
	// 	}
	// }()

	// if err := filepath.WalkDir("./", func(path string, d fs.DirEntry, err error) error {
	// 	if err != nil {
	// 		return err
	// 	}

	// 	if d.IsDir() && !strings.Contains(path, ".git") && !strings.Contains(path, "dist") {
	// 		log.Infof("adding path to watcher: %s", path)
	// 		return watcher.Add(path)
	// 	}

	// 	return nil
	// }); err != nil {
	// 	log.Fatalf("Filepath walk failed: %s", err.Error())
	// 	return
	// }

	// err = watcher.Add("./")
	// if err != nil {
	// 	log.Fatalf("Add failed: %s", err.Error())
	// }

	// <-done
}
