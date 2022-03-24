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
	btrWatch()
}

func btrWatch() {
	log.Println("main")
	log.Println("starting battery watcher")
	for {
		b, err := battery.GetAll()
		if err != nil {
			log.Panic(err)
		}
		btrlvl := math.Round(b[0].Current / b[0].Full * 100)
		state := strings.ToLower(b[0].State.String())

		// log.Printf("b[0].State ---------- %v", b[0].State)
		// log.Printf("b[0].Current -------- %v", b[0].Current)
		// log.Printf("b[0].Full ----------- %v", b[0].Full)
		// log.Printf("b[0].Design --------- %v", b[0].Design)
		// log.Printf("b[0].ChargeRate ----- %v", b[0].ChargeRate)
		// log.Printf("b[0].Voltage -------- %v", b[0].Voltage)
		// log.Printf("b[0].DesignVoltage -- %v", b[0].DesignVoltage)

		var notify = notificator.New(notificator.Options{
			AppName: "Battery Alert",
		})
		if state == "discharging" && btrlvl < minThreshHold {
			notify.Push("Battery Alert", fmt.Sprintf("battery has reached %0.f%s", btrlvl, "%"), "", notificator.UR_CRITICAL)
			writeToFile(fmt.Sprintf("%s battery level alert: %0.f%s\n", logPrepend(), btrlvl, "%"))
		}
		if state != "discharging" && btrlvl > maxThreshHold {
			notify.Push("Battery Alert", fmt.Sprintf("battery has reached %0.f%s", btrlvl, "%"), "", notificator.UR_CRITICAL)
			writeToFile(fmt.Sprintf("%s battery level alert: %0.f%s\n", logPrepend(), btrlvl, "%"))
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
