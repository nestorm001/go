package main

import (
	//    "fmt"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	// t := time.Now().Unix()
	// fmt.Println(t)

	// fmt.Println(time.Unix(t, 0).String())

	// t = time.Now().UnixNano()
	// fmt.Println(t)
	// fmt.Println("------------------")

	// fmt.Println(time.Now().String())

	// fmt.Println(time.Now().Format("2006.01.02 15:04:05"))

	now := time.Now()
	// year,mon,day := now.UTC().Date()
	// hour,min,sec := now.UTC().Clock()
	// zone,_ := now.UTC().Zone()
	// fmt.Printf("UTC time is %d-%d-%d %02d:%02d:%02d %s\n",
	// year,mon,day,hour,min,sec,zone)

	// year,mon,day := now.Date()
	hour, min, sec := now.Clock()
	// zone,_ := now.Zone()
	// fmt.Printf("local time is %d-%d-%d %02d:%02d:%02d %s\n", year,mon,day,hour,min,sec,zone)

	nowSeconds := hour*3600 + min*60 + sec
	shutdownSeconds := 17*3600 + 35*60
	timeToShutdown := shutdownSeconds - nowSeconds
	if timeToShutdown < 0 {
		cmd := exec.Command("cmd.exe", "/c", "shutdown -f")
		cmd.Run()
	} else {
		cmd := exec.Command("cmd.exe", "/c", "shutdown -s -t", strconv.Itoa(timeToShutdown))
		cmd.Run()
	}
}
