//__author__ = "YaoYao"
//Date: 2019-07-19
package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"io/ioutil"
	"mydocker/pkg/command"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const usage = `mydocker`

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage
	app.Commands = command.Commands
	app.Before = func(context *cli.Context) error {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetOutput(os.Stdout)
		log.Info("mydocker")
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("app run %s", err)
	}
}

func namespaceDemo() {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNET | syscall.CLONE_NEWUSER | syscall.CLONE_NEWNS | syscall.CLONE_NEWIPC | syscall.CLONE_NEWPID |
			syscall.CLONE_NEWUTS | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      0,
				Size:        1,
			},
		},
		GidMappings: []syscall.SysProcIDMap{
			{
				ContainerID: 1234,
				HostID:      0,
				Size:        1,
			},
		},
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func cgroupDemo() {
	const cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"
	if os.Args[0] == "/proc/self/exe" {
		//容器进程
		fmt.Printf("current pid %d", syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		os.Exit(-1)
	}
	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWIPC |
			syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET,
	}
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(-1)
	} else {
		fmt.Printf("%v", cmd.Process.Pid)
		err := os.Mkdir(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit"), 0755)
		if err != nil {
			fmt.Println("ERROR", err)
			os.Exit(-1)
		}
		err = ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)
		if err != nil {
			fmt.Println("ERROR", err)
			os.Exit(-1)
		}
		err = ioutil.WriteFile(path.Join(cgroupMemoryHierarchyMount, "testmemorylimit", "memory.limit_in_bytes"), []byte("100m"), 0644)
		if err != nil {
			fmt.Println("ERROR", err)
			os.Exit(-1)
		}
		_, err = cmd.Process.Wait()
	}

}
