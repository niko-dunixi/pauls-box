package helper

import (
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"os"
	"os/exec"
	"strings"
)

func SimpleRun(name string, arg ...string) error {
	prefix := aurora.Bold(aurora.White(`⌬`))
	fmt.Printf("%s %s %s\n", prefix, name, strings.Join(arg, " "))
	command := exec.Command(name, arg...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
