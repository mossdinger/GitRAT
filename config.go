package main

// C2
// Repo Url Format: https://{Token}:x-oauth-basic@github.com/{github_username}/repo
var commandRepoUrl string = ""
var outputRepoUrl string = ""

// Delay
var refreshRate int = 20 // Second

// Sustain
var isIgnoringError bool = true // false will exit when error