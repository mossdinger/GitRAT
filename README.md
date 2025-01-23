# GitRAT
GitRAT is a fileless Remote Access Tool using GitHub repository for running commands and storing outputs. The inspiration behind this project stems from the prevalent practice within many organizations to permit connections to any GitHub repository. This project only serves as a Proof-of-Concept that GitHub can be abused; it is by no means a refined project.

### Disclaimer
* GitRAT is strictly for Educational purposes only.
* Should this project ever be abused, the following behaviors may be used for detection
  * Multiple Connections to 2 Github Repositories from a single process
  * Multiple Cmd/shell instances launched from a process connecting to 2 Github Repositories
* Alternatively, extracting the tokens from the executables then saving them in any public repository would disable them.
* Please create an issue on this repository for any question or concern

## Usage
### Setup
1. Create two Repositories in Github: one will be used for pushing command and one will be used for storing the command output.
2. Create two fine-grained Tokens, one for each repository. The command token should have read access to the content of the command repository while the output token should have both read and write access to the content of the output repository.
3. Create empty file `command` in the command repo and `out` in the output repo.
4. Edit the `config.go` file
5. Build the project

### Post Execution
1. edit the `command` file in the command repo to enter the desired command
2. check the `out` file in the output repo for the output of the command if the command has one

