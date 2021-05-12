package cmd

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func Run(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimRight(out.String(), "\n")
}
