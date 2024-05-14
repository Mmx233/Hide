package tools

import (
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
)

func CreateProcess(name string, args ...string) error {
	program := name
	if len(args) != 0 {
		program += " " + strings.Join(args, " ")
	}
	cmdLine, err := windows.UTF16PtrFromString(strings.Replace(program, `/`, `\\`, -1))
	if err != nil {
		return err
	}

	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var procInfo syscall.ProcessInformation
	startupInfo := &syscall.StartupInfo{}
	return syscall.CreateProcess(
		nil, cmdLine,
		nil, nil, false, windows.CREATE_NO_WINDOW, nil,
		windows.StringToUTF16Ptr(pwd), startupInfo, &procInfo)
}
