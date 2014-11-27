package commands

import (
	"fmt"
)

type VersionCommand struct {
	AppName string
	Version string
}

func (c *VersionCommand) Help() string {
	return fmt.Sprintf("Help %s/%s\n", c.AppName, c.Version)
}

func (c *VersionCommand) Run(args []string) int {
	fmt.Printf("Run %s/%s\n", c.AppName, c.Version)
	return 1
}

func (c *VersionCommand) Synopsis() string {
	return fmt.Sprintf("Synopsis %s/%s", c.AppName, c.Version)
}
