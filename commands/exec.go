package commands

import (
	"flag"
	"fmt"
	"strings"
)

type ExecCommand struct {
	Cmd string
}

func (c *ExecCommand) Help() string {
	helpText := `
Usage: 
`

	return strings.TrimSpace(helpText)
}

func (c *ExecCommand) Run(args []string) int {
	var (
		boolFlag   bool
		intFlag    int
		stringFlag string
	)
	cmdFlags := flag.NewFlagSet("build", flag.ContinueOnError)
	cmdFlags.Usage = func() { c.Help() }
	cmdFlags.BoolVar(&boolFlag, "bool", true, "bool flag")
	noArgsBool := cmdFlags.Bool("b", false, "bool flag")
	cmdFlags.IntVar(&intFlag, "int", 0, "int flag")
	cmdFlags.IntVar(&intFlag, "i", 0, "int flag")
	cmdFlags.StringVar(&stringFlag, "string", "blank", "string flag")
	cmdFlags.StringVar(&stringFlag, "s", "blank", "string flag")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}
	fmt.Println("no args bool flag   =", *noArgsBool)

	fmt.Println("bool flag   =", boolFlag)
	fmt.Println("int flag    =", intFlag)
	fmt.Println("string flag =", stringFlag)

	fmt.Printf("Run %s %+v\n", c.Cmd, args)
	return 1
}

func (c *ExecCommand) Synopsis() string {
	return fmt.Sprintf("Synopsis %s", c.Cmd)
}
