package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	var projectDir, projectName, GIT_REPO_LINK string
	fmt.Println("Enter Project directory:")
	fmt.Scan(&projectDir)
	fmt.Println("Enter Project name:")
	fmt.Scan(&projectName)
	fmt.Println("Enter git origin:")
	fmt.Scan(&GIT_REPO_LINK)
	var DIRECTORY = projectDir + projectName

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
