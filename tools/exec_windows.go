package tools

import (
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
)

func CreateProcess(args ...string) error {
	var program = strings.ReplaceAll(args[0], `\`, `\\`)
	for _, arg := range args[1:] {
		arg = strings.ReplaceAll(arg, `\`, `\\`)
		if strings.Contains(arg, " ") {
			arg = `"` + arg + `"`
		}
		program = program + " " + arg
	}

	cmdLine, err := windows.UTF16PtrFromString(program)
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
