package main

import (
	"log"
	"os"
	"os/exec"
)

func main() {
	err := os.Mkdir("/home/oliver/MyGitHub/"+os.Args[1], 0755)
	if err != nil {
		log.Fatal(err)
	}
	// exec.Command("cd", "/home/oliver/MyGitHub/projectSetUpToolGolang/"+os.Args[1])
	// err2 := exec.Command("touch", "README.md")
	// if err2 != nil {
	// 	log.Fatal(err)
	// }
	os.Chdir("/home/oliver/MyGitHub/" + os.Args[1])
	// cmd := exec.Command("cd", "/home/oliver/MyGitHub/projectSetUpToolGolang/"+os.Args[1])
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
	exec.Command("touch", "README.md").Run()
	exec.Command("git", "init").Run()
}
