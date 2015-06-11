package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	fileName := "monkey.txt"
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	monkey, err := os.Create(fileName)
	if err != nil {
		monkey, _ := os.Open(fileName)
		monkey.WriteString(currentTime + "\n")
	} else {
		monkey.WriteString(currentTime + "\n")
	}
	defer monkey.Close()

	cmd := exec.Command("git", "add", ".")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("git", "commit", "-m", "\"for the monkey\"")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("git", "push")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
