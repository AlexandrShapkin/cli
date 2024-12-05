package cli

import "fmt"

// Cli command processer
type Cli interface {
	OneCmd(input string) error // Process one command
	AddCmd(commands ...*Command) // Adds one or more commands
}

type cli struct {
	cmds map[string]*Command
	parser CommandParser
}

func NewCli() Cli {
	return &cli{
		cmds: make(map[string]*Command),
		parser: NewCommandParser(),
	}
}

// Process one command of the form <command> <flags> <args>
func (c cli) OneCmd(input string) error {
	parsed, err := c.parser.ParseCommand(input)
	if err != nil {
		return err
	}
	cmd, ok := c.cmds[parsed.Name]
	if !ok {
		return fmt.Errorf("command %s not found", input)
	}
	flags := make(map[string]*ParsedCommandFlags)
	for i := range parsed.Flags {
		flag := cmd.GetFlag(i)
		if flag == nil {
			continue
		}
		t := flag.Type
		flags[t] = &ParsedCommandFlags{
			Type: t,
			Args: parsed.Flags[i].Args,
			Name: parsed.Flags[i].Name,
		}
	}
	cmd.Run(flags, parsed.Args)
	return nil
}

// Adds one or more commands
func (c cli) AddCmd(commands ...*Command) {
	for _, cmd := range commands {
		c.cmds[cmd.Use] = cmd
	}
}