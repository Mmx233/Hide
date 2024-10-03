package main

import (
	"fmt"
	"github.com/Mmx233/Hide/tools"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("args is required")
	}
	err := tools.CreateProcess(os.Args[1:]...)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("process created successfully")
	}
}
