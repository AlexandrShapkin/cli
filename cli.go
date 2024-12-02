package cli

import "fmt"

type Cli interface {
	OneCmd(input string) error
	AddCmd(commands ...*Command)
}

type cli struct {
	cmds map[string]*Command
}

func NewCli() Cli {
	return &cli{
		cmds: make(map[string]*Command),
	}
}

func (c cli) OneCmd(input string) error {
	cmd, ok := c.cmds[input]
	if !ok {
		return fmt.Errorf("command %s not found", input)
	}
	cmd.Run()
	return nil
}

func (c cli) AddCmd(commands ...*Command) {
	for _, cmd := range commands {
		c.cmds[cmd.Use] = cmd
	}
}