package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"

	sigar "github.com/cloudfoundry/gosigar"
	"github.com/gen2brain/beeep"
)

const path = "/home/hugo/logs"

var (
	w sync.WaitGroup
)

func main() {
	for true {
		rate := memRate()
		if rate > float64(0.90) {
			shell(fmt.Sprintf("kill %s", strings.ReplaceAll(shell("pidof firefox"), "\n", "")))
			WriteToFile(fmt.Sprintf("%s/%s", path, "memlog"), fmt.Sprintf("%s: %s\n", logPrepend(), "Max rate threshold reached; killing firefox"))
			alert()
			time.Sleep(time.Duration(10) * time.Second)
		}
		time.Sleep(time.Duration(10) * time.Second)
	}
}

func shell(arg string) string {
	args := strings.Split(arg, " ")
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		log.Printf("Command: '%v';\nOutput: %v;\nError: %v\n", arg, string(out), err)
		return err.Error()
	}
	return string(out)
}

func WriteToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func logPrepend() string {
	t := time.Now().Format("2006-01-02 15:04:05")
	output := strings.ReplaceAll(t, "-", "")
	output = strings.ReplaceAll(output, ":", "")
	output = strings.ReplaceAll(output, " ", "-")
	return output
}

func alert() {
	err := beeep.Alert("High memory usage", "Killing firefox", "")
	if err != nil {
		log.Println(err)
	}
}

func memRate() float64 {
	mem := sigar.Mem{}
	mem.Get()
	rate := float64(mem.ActualUsed) / float64(mem.Total)
	WriteToFile(fmt.Sprintf("%s/%s", path, "memlog"), fmt.Sprintf("%s: %.5f\n", logPrepend(), rate))
	return float64(rate)
}
