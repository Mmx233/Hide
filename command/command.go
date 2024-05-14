package command

import (
	"github.com/alecthomas/kingpin/v2"
	"os"
)

func New() (*Command, error) {
	var commands Command

	app := kingpin.New("hide", "Run hidden terminal on windows.")
	app.HelpFlag.Short('h')

	app.Flag("create-process", "Create new process instead of run directly.").BoolVar(&commands.CreateProcess)
	app.Arg("name", "Program to run.").Required().StringVar(&commands.Name)
	app.Arg("args", "Program args.").StringsVar(&commands.Args)

	_, err := app.Parse(os.Args[1:])
	return &commands, err
}

type Command struct {
	CreateProcess bool

	Name string
	Args []string
}
