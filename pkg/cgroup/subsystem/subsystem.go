//__author__ = "YaoYao"
//Date: 2019-07-22
package subsystem

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
)

type ResourceConfig struct {
	MemoryList string
	CpuShare   string
	CpuSet     string
}

type SubSystem interface {
	Name() string
	Set(path string, res *ResourceConfig) error
	Apply(path string, pid int) error
	Remove(path string) error
}

var (
	SubSystemInc = []SubSystem{
		&CpuSetSubSystem{},
		&MemorySubSystem{},
		&CpuSubSyStem{},
	}
)

type MemorySubSystem struct {
}

func (s *MemorySubSystem) Set(cgroupPath string, res *ResourceConfig) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if res.MemoryList != "" {
			if err := ioutil.WriteFile(path.Join(subsysCgroupPath, "memory.limit_in_bytes"), []byte(res.MemoryList), 0644); err != nil {
				return fmt.Errorf("set cgroup memory fail %v", err)
			}
		}
		return nil
	} else {
		return err
	}
}

func (s *MemorySubSystem) Apply(cgroupPath string, pid int) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		if err = ioutil.WriteFile(path.Join(subsysCgroupPath, "tasks"), []byte(strconv.Itoa(pid)), 0644); err != nil {
			return fmt.Errorf("set cgroup proc fail %v", err)
		}
		return nil
	} else {
		return err
	}
}

func (s *MemorySubSystem) Remove(cgroupPath string) error {
	if subsysCgroupPath, err := GetCgroupPath(s.Name(), cgroupPath, true); err == nil {
		return os.Remove(subsysCgroupPath)
	} else {
		return err
	}
}

func (s *MemorySubSystem) Name() string {
	return "memory"
}

type CpuSetSubSystem struct {
}

type CpuSubSyStem struct {
}

func FindCgroupMountPoint(subsystem string) string {
	f, err := os.Open("/proc/self/mountinfo")
	if err != nil {
		return ""
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		txt := scanner.Text()
		fields := strings.Split(txt, " ")
		for _, opt := range strings.Split(fields[len(fields)-1], "") {
			if opt == subsystem {
				return fields[4]
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return ""
	}
	return ""
}

func GetCgroupPath(subsystem string, cgroupPath string, autoCreate bool) (string, error) {
	cgroupRoot := FindCgroupMountPoint(subsystem)
	if _, err := os.Stat(path.Join(cgroupRoot, cgroupPath)); err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), 0755); err == nil {

			} else {
				return "", fmt.Errorf("error create group %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil
	} else {
		return "", fmt.Errorf("cgroup ptah error %v", err)
	}
}
