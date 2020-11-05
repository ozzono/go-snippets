package main

import (
	"fmt"
	"os"
)

func main() {
	var attr = os.ProcAttr{
		Dir: ".",
		Env: os.Environ(),
		Files: []*os.File{
			os.Stdin,
			nil,
			nil,
		},
	}
	process, err := os.StartProcess("/home/hugo/Android/Sdk/emulator/emulator", []string{"/home/hugo/Android/Sdk/emulator/emulator", "-avd", "lite"}, &attr)
	if err == nil {

		// It is not clear from docs, but Realease actually detaches the process
		err = process.Release()
		if err != nil {
			fmt.Println(err.Error())
		}

	} else {
		fmt.Println(err.Error())
	}
}
