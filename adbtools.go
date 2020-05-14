package adbtools

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"time"
)

var (
	deviceID string
	loglvl   bool
)

// TODO: Validate the need of the given commands
// free memory verification and storage: adb shell cat /proc/meminfo |grep MemFree

// Shell executes the given command in the Linux bash terminal
// and return the command output as string
func Shell(arg string) string {
	if len(deviceID) > 0 {
		arg = strings.Replace(arg, "adb", fmt.Sprintf("adb -s %s", deviceID), -1)
	}
	if loglvl {
		log.Println("adb shell: " + arg)
	}
	args := strings.Split(arg, " ")
	out, err := exec.Command(args[0], args[1:]...).Output()
	if err != nil {
		log.Printf("Command: '%v'; Output: %v; Error: %v\n", arg, string(out), err)
		return err.Error()
	}
	if out != nil && len(out) > 0 {
		return fmt.Sprintf("Output:\n %s", out)
	}
	return string(out)
}

// IsActive verifies if the given package is on foreground
func IsActive(appPackage string) bool {
	// TODO: futurually add string normalization
	if strings.Contains(strings.ToLower(Shell("adb shell dumpsys window windows|grep Focus")), strings.ToLower(appPackage)) {
		return true
	}
	return false
}

// TapScreen taps the given coords and waits the given delay in Milliseconds
func TapScreen(x, y, delay int) {
	Shell(fmt.Sprintf("adb shell input tap %d %d", x, y))
	sleep(delay)
	return
}

//sets a sleep wait time in Millisecond
func sleep(delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

// XMLScreen fetches the screen xml data
func XMLScreen(newdump bool) string {
	if newdump {
		Shell("adb shell uiautomator dump")
	}
	return Shell("adb shell cat /sdcard/window_dump.xml")
}

// TapCleanInput tap and cleans the input
func TapCleanInput(x, y, charcount int) {
	charcount = charcount/2 + 1
	TapScreen(x, y, 0)
	Shell("adb shell input keyevent KEYCODE_MOVE_END")
	for i := 0; i < charcount; i++ {
		Shell(`adb shell input keyevent --longpress $(printf 'KEYCODE_DEL %.0s' {1..2})`)
	}
}

func Swipe(coords [4]int) {
	Shell(fmt.Sprintf("adb shell input swipe %d %d %d %d", coords[0], coords[1], coords[2], coords[3]))
}

func CloseApp(app string) {
	Shell(fmt.Sprintf("adb shell am force-stop %s", app))
}

// ClearApp clears all the app data
func ClearApp(app string) error {
	output := Shell(fmt.Sprintf("adb shell pm clear %s", app))
	if strings.Contains(output, "Success") {
		return nil
	}
	return fmt.Errorf("Failed to clear %s app data. Output: %s", app, output)
}

// DeviceID sets the device Id to be used by the adb
func DeviceID(dID string) {
	deviceID = dID
}

// Loglvl enables the logging of every shell command
func Loglvl(verbose bool) {
	loglvl = verbose
}

func InputText(text string, splitted bool) error {
	if len(text) == 0 {
		return fmt.Errorf("invalid input; cannot be empty")
	}
	// Fixes whitespace input with adb and shell
	text = strings.Replace(text, " ", "\\s", -1)
	if splitted {
		for _, textRune := range strings.Split(text, "") {
			Shell(fmt.Sprintf("adb shell input text %v", textRune))
		}
		return nil
	}
	Shell("adb shell input text %s" + text)
	return nil
}

func PageDown() {
	// code 93 is equivalent to "KEYCODE_PAGE_DOWN"
	Shell("adb shell input keyevent 93")
}

func PageUp() {
	// code 92 is equivalent to "KEYCODE_PAGE_UP"
	Shell("adb shell input keyevent 92")
}
