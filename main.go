package main

import "time"

func handleError(err error, isIgnoreError bool) {
	if err != nil && !isIgnoreError {
		print(err)
		return
	}
}

func loopCommand(config *Config) {

	for {
		commandString, err := LoadCommand(config.CommandRepoUrl)
		handleError(err, config.IsIgnoringError)

		commandOutput, err := ExecuteCommand(commandString, config.Timeout)
		handleError(err, config.IsIgnoringError)

		err = PushOutRepo(commandOutput, config.OutputRepoUrl)
		handleError(err, config.IsIgnoringError)

		time.Sleep(time.Duration(config.RefreshInterval) * time.Second)
	}
}

func main() {
	config := LoadConfig()
	loopCommand(config)
}
