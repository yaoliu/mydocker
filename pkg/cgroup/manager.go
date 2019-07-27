//__author__ = "YaoYao"
//Date: 2019-07-21
package cgroup

import (
	log "github.com/sirupsen/logrus"
	"mydocker/pkg/cgroup/subsystem"
)

type Manager struct {
	Path     string
	Resource *subsystem.ResourceConfig
}

func NewCgroupManager(path string) *Manager {
	return &Manager{
		Path: path,
	}
}

func (c *Manager) Apply(pid int) error {
	for _, subSysIns := range subsystem.SubSystemInc {
		subSysIns.Apply(c.Path, pid)
	}
	return nil
}

func (c *Manager) Set(res *subsystem.ResourceConfig) error {
	for _, subSysIns := range subsystem.SubSystemInc {
		subSysIns.Set(c.Path, res)
	}
	return nil
}

func (c *Manager) Destroy() error {
	for _, subSysIns := range subsystem.SubSystemInc {
		if err := subSysIns.Remove(c.Path); err != nil {
			log.Error(err)
		}
	}
	return nil
}
