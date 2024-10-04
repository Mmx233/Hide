package tools

import (
	"golang.org/x/sys/windows"
	"os"
	"os/exec"
	"syscall"
)

func CreateProcess(name string, args ...string) error {
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd := exec.Command(name, args...)
	cmd.Dir = pwd
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:       true,
		CreationFlags:    windows.CREATE_NEW_CONSOLE | windows.CREATE_NO_WINDOW,
		NoInheritHandles: true,
	}
	return cmd.Start()
}
