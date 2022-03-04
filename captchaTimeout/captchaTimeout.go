package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	defaultSleepTime = 100
)

func main() {
	shell("adb shell am force-stop org.mozilla.firefox")
	start := time.Now().UnixNano()
	count := 0
	for  {
		count++
		fmt.Printf("Try count: %d\n", count)
		sleeper(100 * count)
	}
	end := time.Now().UnixNano()
	total := end - start
	if total > 1000000000 {
		fmt.Printf("Total time: %v\n", total/1000000000)
	} else {
		fmt.Printf("verification total time: %v\n", time.Duration(total)*time.Millisecond)
	}
	fmt.Printf("count: %d\n", count)
}

func sleeper(delay int) {
	fmt.Printf("sleep: %vs\n", delay*defaultSleepTime/1000)
	for i := 0; i < delay; i++ {
		time.Sleep(time.Duration(defaultSleepTime) * time.Millisecond)
	}
}

func captchaTimeout() bool {
	shell("adb shell am start -a android.activity.MAIN -n org.mozilla.firefox/org.mozilla.gecko.BrowserApp")
	sleeper(10)
	shell("adb shell am force-stop org.mozilla.firefox")
	shell("adb shell am start -a android.activity.MAIN -n org.mozilla.firefox/org.mozilla.gecko.BrowserApp -d https://minhaconta.riocardmais.com.br/cadastro/home --ez private_tab true")
	sleeper(50)
	shell("adb shell uiautomator dump")
	xmlscreen := shell("adb shell cat /sdcard/window_dump.xml")
	shell("adb shell pm clear org.mozilla.firefox")
	return strings.Contains(strings.ToLower(xmlscreen), "additional security check")
}

func shell(arg string) string {
	args := strings.Split(arg, " ")
	if len(args) == 1 {
		args = append(args, "")
	} else if len(args) < 1 {
		log.Println("Invalid command")
		return ""
	}
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		log.Printf("Command: '%v'; Output: %v; Error: %v\n", arg, string(out), err)
		return err.Error()
	}
	return string(out)
}
