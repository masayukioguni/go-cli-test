package commands

import (
	"fmt"
)

type ExecCommand struct {
	Cmd string
}

func (c *ExecCommand) Help() string {
	return fmt.Sprintf("Help %s\n", c.Cmd)
}

func (c *ExecCommand) Run(args []string) int {
	fmt.Printf("Run %s\n", c.Cmd)
	return 1
}

func (c *ExecCommand) Synopsis() string {
	return fmt.Sprintf("Synopsis %s", c.Cmd)
}
