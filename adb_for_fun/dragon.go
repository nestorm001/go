package main

import (
	"os/exec"
	"fmt"
	"strings"
	"time"
)

func main() {
	for true {
		topActivity := getTopActivity()
		if strings.Contains(topActivity, "dragon") {
			performClick()
		} else {
			fmt.Println(topActivity)
			time.Sleep(10 * time.Second)
		}
	}
}

func getTopActivity() string {
	cmd := exec.Command("adb", "shell", "dumpsys activity | grep top-activity")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func performClick() {
	for i := 0; i < 20; i++ {
		click()
		time.Sleep(100 * time.Millisecond)
	}
}

func click() {
	cmd := exec.Command("adb", "shell", "input tap 700 1250")
	cmd.CombinedOutput()
}
