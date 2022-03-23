package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/0xAX/notificator"
	"github.com/distatus/battery"
)

const (
	minThreshHold = 20
	maxThreshHold = 90
)

var logPath string

func init() {
	flag.StringVar(&logPath, "log-path", "", "log file path")
}

func main() {
	flag.Parse()
	b, err := battery.GetAll()
	if err != nil {
		log.Panic(err)
	}
	btrWatch(*b[0])
}

func btrWatch(b battery.Battery) {
	log.Println("main")
	log.Println("starting battery watcher")
	for {
		btrlvl := math.Round(b.Current / b.Full * 100)
		state := strings.ToLower(b.State.String())

		var notify = notificator.New(notificator.Options{
			AppName: "Battery Alert",
		})
		if state == "discharging" && btrlvl < minThreshHold {
			notify.Push("Battery Alert", fmt.Sprintf("Discharding battery has reached %f%s", btrlvl, "%"), "", notificator.UR_CRITICAL)
			writeToFile(fmt.Sprintf("%s battery level alert: %f%s\n", logPrepend(), btrlvl, "%"))
		}
		if state == "charging" && btrlvl > maxThreshHold {
			notify.Push("Battery Alert", fmt.Sprintf("Discharding battery has reached %f%s", btrlvl, "%"), "", notificator.UR_CRITICAL)
			writeToFile(fmt.Sprintf("%s battery level alert: %f%s\n", logPrepend(), btrlvl, "%"))
		}
		time.Sleep(time.Minute)
	}
}

func logPrepend() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func writeToFile(data string) error {
	if logPath == "" {
		return nil
	}
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
