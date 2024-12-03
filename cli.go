package cli

import "fmt"

// Cli command processer
type Cli interface {
	OneCmd(input string) error // Process one command
	AddCmd(commands ...*Command) // Adds one or more commands
}

type cli struct {
	cmds map[string]*Command
}

func NewCli() Cli {
	return &cli{
		cmds: make(map[string]*Command),
	}
}

// Process one command of the form <command> <flags> <args>
func (c cli) OneCmd(input string) error {
	cmd, ok := c.cmds[input]
	if !ok {
		return fmt.Errorf("command %s not found", input)
	}
	cmd.Run()
	return nil
}

// Adds one or more commands
func (c cli) AddCmd(commands ...*Command) {
	for _, cmd := range commands {
		c.cmds[cmd.Use] = cmd
	}
}