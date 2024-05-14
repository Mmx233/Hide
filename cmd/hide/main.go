package main

import (
	"github.com/Mmx233/Hide/command"
	"github.com/Mmx233/Hide/tools"
	"log"
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
		err := exec.Command(config.Name, config.Args...).Run()
		if err != nil {
			log.Fatalln(err)
		}
	}
}
