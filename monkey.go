package main

import (
	"os"
	"os/exec"
	"time"
)

func main() {
	fileName := "monkey.txt"
	currentTime := time.Now().Format("2006.01.02 15:04:05")
	monkey, err := os.OpenFile(fileName, os.O_APPEND, 0)
	if err != nil {
		monkey, _ := os.Create(fileName)
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
	cmd = exec.Command("git", "nestorm001")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
