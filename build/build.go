package main

import (
	"os"

	"github.com/roemer/gotaskr"
)

func main() {
	os.Exit(gotaskr.Execute())
}

func init() {
	gotaskr.Task("test", func() error {
		// This is a placeholder for the test task.
		// You can add your test logic here.
		return nil
	})
}
