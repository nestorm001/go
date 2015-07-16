package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "python", "-m", "pydoc", "-p", "23333")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
