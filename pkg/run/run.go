//__author__ = "YaoYao"
//Date: 2019-07-21
package run

import (
	log "github.com/sirupsen/logrus"
	"mydocker/pkg/container"
	"os"
)

func Run(tty bool, command string) {
	parent := container.NewParentProcess(tty, command)
	if err := parent.Start(); err != nil {
		log.Error("Run", err)
	}
	_ = parent.Wait()
	os.Exit(-1)
}
