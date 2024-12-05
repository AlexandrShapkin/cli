package cli

import (
	"strings"
)

// Description of command or flag or something else
type Description struct {
	Long  string
	Short string
}

// Command structure representing a command
type Command struct {
	Use   string
	Flags []*CommandFlag
	Desc  Description
	Run   func(flags map[string]*ParsedCommandFlags, args []string)
}

// Returns a flag by type name. Returns a flag by type name. If it does not exist, then nil.
func (c *Command) GetFlag(flag string) *CommandFlag {
	if c.Flags == nil {
		return nil
	}
	for _, f := range c.Flags {
		long := strings.Compare(f.Long, flag) == 0
		short := strings.Compare(f.Short, flag) == 0
		if long || short {
			return f
		}
	}
	return nil
}

// CommandFlag structure representing flag for command
type CommandFlag struct {
	Type  string // flag id
	Long  string // for '--flag'
	Short string // for '-f'
	Desc  Description
}

type ParsedCommand struct {
	Name  string // same as Use in Command
	Flags map[string]*ParsedCommandFlags
	Args  []string
}

type ParsedCommandFlags struct {
	Type string
	Name string
	Args string
}
