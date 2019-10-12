package executil

import (
	"os"
	"os/exec"
	"sync"

	"github.com/sirupsen/logrus"
)

func RunAsync(wg *sync.WaitGroup, name string, arg ...string) {
	wg.Add(1)
	go func() {
		cmd := exec.Command(name, arg...)
		// var out bytes.Buffer
		cmd.Stdout = os.Stdout
		err := cmd.Run()
		if err != nil {
			logrus.Fatal(err)
		}
		wg.Done()
	}()
}

func Run(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	// var out bytes.Buffer
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}
