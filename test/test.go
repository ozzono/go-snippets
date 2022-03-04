package main
import (
	"fmt"
	"strings"
	// "log"
	"os"
	"os/exec"
)

func main(){

	// cmd:="ls"
	args:=os.Args[1:]
	arg:=strings.Join(args," ")

	// output,_ := exec.Command(string(args[0])).Output()
	output,_ := exec.Command(arg).Output()
	fmt.Printf("Command: %s\n",arg)
	fmt.Printf(string(output))
}