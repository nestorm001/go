package main

import (
	"fmt"
	"os/exec"
	"strings"
	"regexp"
)

const OVERRIDE = "ANDROID_DAILY_OVERRIDE"
const IDE_PATH = "H:\\ide\\android-studio\\bin\\studio64.exe"

func sync() string {
	cmd := exec.Command("gradlew")

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return string(out)
}

func getDailyHash(syncResult string) string {
	pat := "\".*\""
	regex := regexp.MustCompile(pat)
	var result string
	for _, s := range strings.Split(syncResult, "\n") {
		if strings.Contains(s, OVERRIDE) {
			result = string(regex.Find([]byte(s))[:])
			result = strings.Trim(result, "\"")
			break
		}
	}
	return result
}

func setDailyOverride(dailyHash string) {
	cmd := exec.Command("setx", OVERRIDE, dailyHash)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}

func restartAS() {
	cmd := exec.Command("taskkill", "/IM", "studio64.exe")
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
	}
	//time.Sleep(3 * time.Second)
	//anotherCmd := exec.Command(IDE_PATH)
	//err = anotherCmd.Start()
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func main() {
	syncResult := sync()
	dailyHash := getDailyHash(syncResult)
	if len(dailyHash) != 0 {
		setDailyOverride(dailyHash)
		restartAS()
	}
}
