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

type device struct {
	ID string
}

// TODO: Validate the need of the given commands
// free memory verification and storage: adb shell cat /proc/meminfo |grep MemFree

// Shell executes the given command in the Linux bash terminal
// and return the command output as string
func (device *device) Shell(arg string) string {
	if len(device.ID) > 0 {
		arg = strings.Replace(arg, "adb", fmt.Sprintf("adb -s %s", device.ID), -1)
	}
	if loglvl {
		log.Println(arg)
	}
	return shell(arg)
}

func shell(arg string) string {
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

// Foreground verifies if the given package is on foreground
func (device *device) Foreground(appPackage string) bool {
	// TODO: futurally add string normalization
	return strings.Contains(strings.ToLower(device.Shell("adb shell dumpsys window windows|grep Focus")), strings.ToLower(appPackage))
}

// TapScreen taps the given coords and waits the given delay in Milliseconds
func (device *device) TapScreen(x, y, delay int) {
	device.Shell(fmt.Sprintf("adb shell input tap %d %d", x, y))
	sleep(delay)
	return
}

//sets a sleep wait time in Millisecond
func sleep(delay int) {
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

// XMLScreen fetches the screen xml data
func (device *device) XMLScreen(newdump bool) string {
	if newdump {
		device.Shell("adb shell uiautomator dump")
	}
	return device.Shell("adb shell cat /sdcard/window_dump.xml")
}

// TapCleanInput tap and cleans the input
func (device *device) TapCleanInput(x, y, charcount int) {
	charcount = charcount/2 + 1
	devices.TapScreen(x, y, 0)
	device.Shell("adb shell input keyevent KEYCODE_MOVE_END")
	for i := 0; i < charcount; i++ {
		device.Shell(`adb shell input keyevent --longpress $(printf 'KEYCODE_DEL %.0s' {1..2})`)
	}
}

func (device *device) Swipe(coords [4]int) {
	device.Shell(fmt.Sprintf("adb shell input swipe %d %d %d %d", coords[0], coords[1], coords[2], coords[3]))
}

func (device *device) CloseApp(app string) {
	device.Shell(fmt.Sprintf("adb shell am force-stop %s", app))
}

// ClearApp clears all the app data
func (device *device) ClearApp(app string) error {
	output := device.Shell(fmt.Sprintf("adb shell pm clear %s", app))
	if strings.Contains(output, "Success") {
		return nil
	}
	return fmt.Errorf("Failed to clear %s app data. Output: %s", app, output)
}

// Loglvl enables the logging of every shell command
func Loglvl(verbose bool) {
	loglvl = verbose
}

func (device *device) InputText(text string, splitted bool) error {
	if len(text) == 0 {
		return fmt.Errorf("invalid input; cannot be empty")
	}
	// Fixes whitespace input with adb and shell
	text = strings.Replace(text, " ", "\\s", -1)
	if splitted {
		for _, textRune := range strings.Split(text, "") {
			device.Shell(fmt.Sprintf("adb shell input text %v", textRune))
		}
		return nil
	}
	device.Shell("adb shell input text %s" + text)
	return nil
}

func (device *device) PageDown() {
	// code 93 is equivalent to "KEYCODE_PAGE_DOWN"
	device.Shell("adb shell input keyevent 93")
}

func (device *device) PageUp() {
	// code 92 is equivalent to "KEYCODE_PAGE_UP"
	device.Shell("adb shell input keyevent 92")
}

func Devices() []device {
	output := []device{}
	for _, row := range strings.Split(device.Shell("adb devices"), "\n") {
		if strings.Contains(row, "device") {
			output = append(output, Device{ID: strings.Split(row, "	")[0]})
		}
	}
	return output
}

func NewDevice(deviceID string) device {
	return device{ID: deviceID}
}
