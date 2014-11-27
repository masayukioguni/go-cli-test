package main

import (
	"github.com/masayukioguni/go-cli-test/commands"
	"github.com/mitchellh/cli"
	"log"
	"os"
)

func main() {
	c := cli.NewCLI("go-cli-test", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"version": func() (cli.Command, error) {
			return &commands.VersionCommand{
				Version: "9.9.9",
				AppName: "go-cli-test",
			}, nil
		},
		"exec": func() (cli.Command, error) {
			return &commands.ExecCommand{
				Cmd: "exec command",
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
