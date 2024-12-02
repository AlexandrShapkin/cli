package cli

import "strings"

type Description struct {
	Long  string
	Short string
}

type Command struct {
	Use   string
	Flags []*CommandFlag
	Desc  Description
	Run   func()
}

func (c *Command) GetFlag(flag string) *CommandFlag {
	for _, f := range c.Flags {
		long := strings.Compare(f.Long, flag) == 0
		short := strings.Compare(f.Short, flag) == 0
		if long || short {
			return f
		}
	}
	return nil
}

type CommandFlag struct {
	Type  string
	Long  string
	Short string
	Desc  Description
}

type ParsedCommand struct {
	Name  string
	Flags map[string]*ParsedCommandFlags
	Args []string
}

type ParsedCommandFlags struct {
	Name string
	Args string
}
