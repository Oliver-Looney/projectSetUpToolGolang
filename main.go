package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	var DIRECTORY = "/home/oliver/MyGitHub/" + os.Args[1]
	var GIT_REPO_LINK = os.Args[2]

	err := os.Mkdir(DIRECTORY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	os.Chdir(DIRECTORY)

	exec.Command("touch", "README.md").Run()
	exec.Command("git", "init").Run()

	exec.Command("git", "remote", "add", "origin", GIT_REPO_LINK).Run()
	exec.Command("git", "add", ".").Run()
	exec.Command("git", "commit", "-m", "\"initial commit\"").Run()
	exec.Command("git", "push", "-u", "origin", "main").Run()
	exec.Command("code", ".").Run()
}
