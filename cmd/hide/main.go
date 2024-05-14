package main

import (
	"github.com/Mmx233/Hide/command"
	"github.com/Mmx233/Hide/tools"
	"log"
	"os"
	"os/exec"
)

func main() {
	config, err := command.New()
	if err != nil {
		log.Fatalln("parse command arguments failed:", err)
	}

	if config.CreateProcess {
		err = tools.CreateProcess(config.Name, config.Args...)
		if err != nil {
			log.Fatalln(err)
		} else {
			log.Println("process created successfully")
		}
	} else {
		cmd := exec.Command(config.Name, config.Args...)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
