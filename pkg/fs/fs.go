//__author__ = "YaoYao"
//Date: 2019-07-24
package fs

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"syscall"
)

func pivotRoot(root string) error {
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("Mount rootfs to  error: %v", err)
	}
	pivotDir := filepath.Join(root, ".pivot_root")
	if err := os.Mkdir(pivotDir, 0777); err != nil {
		return err
	}
	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return fmt.Errorf("pivot_root %v", err)
	}
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chadir / %v", err)
	}
	pivotDir = filepath.Join("/", ".pivot_root")
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount pivot_root dir %v", err)
	}
	return os.Remove(pivotDir)
}

func setUpMount() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Errorf("Get current location error %v", err)
		return
	}
	log.Infof("Current location is %s", pwd)
	_ := pivotRoot(pwd)
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	_ := syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	_ = syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
}
