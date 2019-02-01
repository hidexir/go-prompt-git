package main

import (
	"os"
	"os/exec"

	prompt "github.com/c-bata/go-prompt"
)

func executor(t string) {
	if t == "add" {
		cmd := exec.Command("git", "add", "-A")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	if t == "commit" {
		cmd := exec.Command("git", "commit")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
	return
}

func completer(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "add"},
		{Text: "commit"},
	}
}

func main() {
	p := prompt.New(
		executor,
		completer,
	)
	p.Run()
}
