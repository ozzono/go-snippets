package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

var (
	folderpath string
	filter     string
)

type fileInfo struct {
	Size int
	Name string
}

type allFiles []fileInfo

func init() {
	flag.StringVar(&folderpath, "p", "./", "Sets the folder to have files sorted")
	flag.StringVar(&filter, "f", "", "Sets the filter of files sortedly renamed")
}

func main() {
	flag.Parse()
	rename(sort(getFiles()))
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

func getFiles() []fileInfo {
	output := []fileInfo{}
	files, err := ioutil.ReadDir(folderpath)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Getting files with filter %s\n", filter)
	for _, f := range files {
		if strings.Contains(f.Name(), filter) {
			output = append(output, fileInfo{int(f.Size()), f.Name()})
		} else {
			fmt.Printf("Skipping %s file\n", f.Name())
		}
	}
	return output
}

func sort(pool []fileInfo) []fileInfo {
	for i := len(pool); i > 0; i-- {
		for j := 1; j < i; j++ {
			if pool[j-1].Size > pool[j].Size {
				tmp := pool[j]
				pool[j] = pool[j-1]
				pool[j-1] = tmp
			}
		}
	}
	return pool
}

func rename(files []fileInfo) {
	path := ""
	if folderpath != "./" {
		path = folderpath
	}
	pad := fmt.Sprintf("%d", len(fmt.Sprintf("%d", len(files))))
	for i := range files {
		mv := fmt.Sprintf("mv %s %s", path+files[i].Name, fmt.Sprintf("%0"+pad+"d_%s", i, path+files[i].Name))
		shell(mv)
	}
}

// This methos was used for undoing the changes made by this script
func unchanged(input string) bool {
	for _, item := range []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"} {
		if item == input {
			return false
		}
	}
	return true
}
