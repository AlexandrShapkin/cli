package cli

import (
	"strings"
	"testing"
)

func Test_cli_OneCmd(t *testing.T) {
	cli := NewCli()
	cli.AddCmd(
		&Command{
			Use: "do",
			Flags: make([]*CommandFlag, 0),
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {},
		},
		&Command{
			Use: "witharg",
			Flags: make([]*CommandFlag, 0),
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				want := "arg"
				if len(args) != 1 || args[0] != want {
					t.Errorf("argument is incorrect")
				}
			},
		},
		&Command{
			Use: "withtwoargs",
			Flags: make([]*CommandFlag, 0),
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				want := []string{"arg0", "arg1"}
				if len(args) != 2 || args[0] != want[0] || args[1] != want[1] {
					t.Errorf("argument is incorrect")
				}
			},
		},
		&Command{
			Use: "withmultiwordarg",
			Flags: make([]*CommandFlag, 0),
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				want := "arg0 arg1"
				if len(args) != 1 || args[0] != want {
					t.Errorf("argument is incorrect")
				}
			},
		},
		&Command{
			Use: "withmultiwordargs",
			Flags: make([]*CommandFlag, 0),
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				want := []string{"arg0 arg1", "arg2 arg3"}
				if len(args) != 2 || args[0] != want[0] || args[1] != want[1] {
					t.Errorf("argument is incorrect")
				}
			},
		},
		&Command{
			Use: "withlongflag",
			Flags: []*CommandFlag {
				{
					Type: "flag",
					Long: "flag",
					Short: "f",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := "flag"
				wantName := "flag"
				flag, ok := flags[wantType]
				if !ok {
					t.Errorf("flag with type %s not found", wantType)
					return
				}
				if strings.Compare(flag.Type, wantType) != 0 {
					t.Errorf("flag type is incorrect want %s, got %s", wantType, flag.Type)
					return
				}
				if strings.Compare(flag.Name, wantName) != 0 {
					t.Errorf("flag name is incorrect want %s, got %s", wantName, flag.Name)
					return
				}
			},
		},
		&Command{
			Use: "withlongargflag",
			Flags: []*CommandFlag {
				{
					Type: "flag",
					Long: "flag",
					Short: "f",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := "flag"
				wantName := "flag"
				wantFlagArg := "flagarg"
				flag, ok := flags[wantType]
				if !ok {
					t.Errorf("flag with type %s not found", wantType)
					return
				}
				if strings.Compare(flag.Type, wantType) != 0 {
					t.Errorf("flag type is incorrect want %s, got %s", wantType, flag.Type)
					return
				}
				if strings.Compare(flag.Name, wantName) != 0 {
					t.Errorf("flag name is incorrect want %s, got %s", wantName, flag.Name)
					return
				}
				if strings.Compare(flag.Args, wantFlagArg) != 0 {
					t.Errorf("flag arg is incorrect want %s, got %s", wantFlagArg, flag.Name)
					return
				}
			},
		},
		&Command{
			Use: "withlongflags",
			Flags: []*CommandFlag {
				{
					Type: "flag0",
					Long: "flag0",
					Short: "f",
				},
				{
					Type: "flag1",
					Long: "flag1",
					Short: "a",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := []string{"flag0", "flag1"}
				wantName := []string{"flag0", "flag1"}
				flag0, ok := flags[wantType[0]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[0])
					return
				}
				flag1, ok := flags[wantType[1]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[1])
					return
				}
				if strings.Compare(flag0.Type, wantType[0]) != 0 || strings.Compare(flag1.Type, wantType[1]) != 0 {
					t.Errorf("flag type is incorrect want %v, got %v", wantType, []string{flag0.Type, flag1.Type})
					return
				}
				if strings.Compare(flag0.Name, wantName[0]) != 0 || strings.Compare(flag1.Name, wantName[1]) != 0 {
					t.Errorf("flag name is incorrect want %v, got %v", wantType, []string{flag0.Name, flag1.Name})
					return
				}
			},
		},
		&Command{
			Use: "withlongflagswithargs",
			Flags: []*CommandFlag {
				{
					Type: "flag0",
					Long: "flag0",
					Short: "f",
				},
				{
					Type: "flag1",
					Long: "flag1",
					Short: "a",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := []string{"flag0", "flag1"}
				wantName := []string{"flag0", "flag1"}
				wantFlagArg := []string{"flagarg0", "flagarg1"}
				flag0, ok := flags[wantType[0]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[0])
					return
				}
				flag1, ok := flags[wantType[1]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[1])
					return
				}
				if strings.Compare(flag0.Type, wantType[0]) != 0 || strings.Compare(flag1.Type, wantType[1]) != 0 {
					t.Errorf("flag type is incorrect want %v, got %v", wantType, []string{flag0.Type, flag1.Type})
					return
				}
				if strings.Compare(flag0.Name, wantName[0]) != 0 || strings.Compare(flag1.Name, wantName[1]) != 0 {
					t.Errorf("flag name is incorrect want %v, got %v", wantType, []string{flag0.Name, flag1.Name})
					return
				}
				if strings.Compare(flag0.Args, wantFlagArg[0]) != 0 || strings.Compare(flag1.Args, wantFlagArg[1]) != 0 {
					t.Errorf("flag args is incorrect want %v, got %v", wantFlagArg, []string{flag0.Args, flag1.Args})
					return
				}
			},
		},
		&Command{
			Use: "withlongflagswithmargs",
			Flags: []*CommandFlag {
				{
					Type: "flag0",
					Long: "flag0",
					Short: "f",
				},
				{
					Type: "flag1",
					Long: "flag1",
					Short: "a",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := []string{"flag0", "flag1"}
				wantName := []string{"flag0", "flag1"}
				wantFlagArg := []string{"flagarg0 flagarg1", "flagarg2 flagarg3"}
				flag0, ok := flags[wantType[0]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[0])
					return
				}
				flag1, ok := flags[wantType[1]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[1])
					return
				}
				if strings.Compare(flag0.Type, wantType[0]) != 0 || strings.Compare(flag1.Type, wantType[1]) != 0 {
					t.Errorf("flag type is incorrect want %v, got %v", wantType, []string{flag0.Type, flag1.Type})
					return
				}
				if strings.Compare(flag0.Name, wantName[0]) != 0 || strings.Compare(flag1.Name, wantName[1]) != 0 {
					t.Errorf("flag name is incorrect want %v, got %v", wantType, []string{flag0.Name, flag1.Name})
					return
				}
				if strings.Compare(flag0.Args, wantFlagArg[0]) != 0 || strings.Compare(flag1.Args, wantFlagArg[1]) != 0 {
					t.Errorf("flag args is incorrect want %v, got %v", wantFlagArg, []string{flag0.Args, flag1.Args})
					return
				}
			},
		},
		&Command{
			Use: "withshortflag",
			Flags: []*CommandFlag {
				{
					Type: "flag",
					Long: "flag",
					Short: "f",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := "flag"
				wantName := "f"
				flag, ok := flags[wantType]
				if !ok {
					t.Errorf("flag with type %s not found", wantType)
					return
				}
				if strings.Compare(flag.Type, wantType) != 0 {
					t.Errorf("flag type is incorrect want %s, got %s", wantType, flag.Type)
					return
				}
				if strings.Compare(flag.Name, wantName) != 0 {
					t.Errorf("flag name is incorrect want %s, got %s", wantName, flag.Name)
					return
				}
			},
		},
		&Command{
			Use: "withshortflags",
			Flags: []*CommandFlag {
				{
					Type: "flag0",
					Long: "flag0",
					Short: "f",
				},
				{
					Type: "flag1",
					Long: "flag1",
					Short: "a",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := []string{"flag0", "flag1"}
				wantName := []string{"f", "a"}
				flag0, ok := flags[wantType[0]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[0])
					return
				}
				flag1, ok := flags[wantType[1]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[1])
					return
				}
				if strings.Compare(flag0.Type, wantType[0]) != 0 || strings.Compare(flag1.Type, wantType[1]) != 0 {
					t.Errorf("flag type is incorrect want %v, got %v", wantType, []string{flag0.Type, flag1.Type})
					return
				}
				if strings.Compare(flag0.Name, wantName[0]) != 0 || strings.Compare(flag1.Name, wantName[1]) != 0 {
					t.Errorf("flag name is incorrect want %v, got %v", wantType, []string{flag0.Name, flag1.Name})
					return
				}
			},
		},
				&Command{
			Use: "withshortflags",
			Flags: []*CommandFlag {
				{
					Type: "flag0",
					Long: "flag0",
					Short: "f",
				},
				{
					Type: "flag1",
					Long: "flag1",
					Short: "a",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := []string{"flag0", "flag1"}
				wantName := []string{"f", "a"}
				flag0, ok := flags[wantType[0]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[0])
					return
				}
				flag1, ok := flags[wantType[1]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[1])
					return
				}
				if strings.Compare(flag0.Type, wantType[0]) != 0 || strings.Compare(flag1.Type, wantType[1]) != 0 {
					t.Errorf("flag type is incorrect want %v, got %v", wantType, []string{flag0.Type, flag1.Type})
					return
				}
				if strings.Compare(flag0.Name, wantName[0]) != 0 || strings.Compare(flag1.Name, wantName[1]) != 0 {
					t.Errorf("flag name is incorrect want %v, got %v", wantType, []string{flag0.Name, flag1.Name})
					return
				}
			},
		},
		&Command{
			Use: "withlongflagandarg",
			Flags: []*CommandFlag {
				{
					Type: "flag",
					Long: "flag",
					Short: "f",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := "flag"
				wantName := "flag"
				flag, ok := flags[wantType]
				if !ok {
					t.Errorf("flag with type %s not found", wantType)
					return
				}
				if strings.Compare(flag.Type, wantType) != 0 {
					t.Errorf("flag type is incorrect want %s, got %s", wantType, flag.Type)
					return
				}
				if strings.Compare(flag.Name, wantName) != 0 {
					t.Errorf("flag name is incorrect want %s, got %s", wantName, flag.Name)
					return
				}
				if len(args) != 1 || args[0] != "arg" {
					t.Errorf("argument is incorrect")
				}
			},
		},
		&Command{
			Use: "withshortflagandarg",
			Flags: []*CommandFlag {
				{
					Type: "flag",
					Long: "flag",
					Short: "f",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := "flag"
				wantName := "f"
				flag, ok := flags[wantType]
				if !ok {
					t.Errorf("flag with type %s not found", wantType)
					return
				}
				if strings.Compare(flag.Type, wantType) != 0 {
					t.Errorf("flag type is incorrect want %s, got %s", wantType, flag.Type)
					return
				}
				if strings.Compare(flag.Name, wantName) != 0 {
					t.Errorf("flag name is incorrect want %s, got %s", wantName, flag.Name)
					return
				}
				if len(args) != 1 || args[0] != "arg" {
					t.Errorf("argument is incorrect")
				}
			},
		},
		&Command{
			Use: "withlongflagandmultiwordargs",
			Flags: []*CommandFlag {
				{
					Type: "flag",
					Long: "flag",
					Short: "f",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := "flag"
				wantName := "flag"
				wantArg := []string{"arg0 arg1", "arg2 arg3"}
				flag, ok := flags[wantType]
				if !ok {
					t.Errorf("flag with type %s not found", wantType)
					return
				}
				if strings.Compare(flag.Type, wantType) != 0 {
					t.Errorf("flag type is incorrect want %s, got %s", wantType, flag.Type)
					return
				}
				if strings.Compare(flag.Name, wantName) != 0 {
					t.Errorf("flag name is incorrect want %s, got %s", wantName, flag.Name)
					return
				}
				if len(args) != 2 || args[0] != wantArg[0] || args[1] != wantArg[1] {
					t.Errorf("argument is incorrect want = %v, got = %v", wantArg, args)
				}
			},
		},
		&Command{
			Use: "withshortflagandmultiwordargs",
			Flags: []*CommandFlag {
				{
					Type: "flag",
					Long: "flag",
					Short: "f",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := "flag"
				wantName := "f"
				wantArg := []string{"arg0 arg1", "arg2 arg3"}
				flag, ok := flags[wantType]
				if !ok {
					t.Errorf("flag with type %s not found", wantType)
					return
				}
				if strings.Compare(flag.Type, wantType) != 0 {
					t.Errorf("flag type is incorrect want %s, got %s", wantType, flag.Type)
					return
				}
				if strings.Compare(flag.Name, wantName) != 0 {
					t.Errorf("flag name is incorrect want %s, got %s", wantName, flag.Name)
					return
				}
				if len(args) != 2 || args[0] != wantArg[0] || args[1] != wantArg[1] {
					t.Errorf("argument is incorrect want = %v, got = %v", wantArg, args)
				}
			},
		},
		&Command{
			Use: "wlfwmaama",
			Flags: []*CommandFlag {
				{
					Type: "flag0",
					Long: "flag0",
					Short: "f",
				},
				{
					Type: "flag1",
					Long: "flag1",
					Short: "a",
				},
			},
			Run: func(flags map[string]*ParsedCommandFlags, args []string) {
				wantType := []string{"flag0", "flag1"}
				wantName := []string{"flag0", "flag1"}
				wantArgs := []string{"arg0 arg1", "arg2 arg3"}
				wantFlagArg := []string{"flagarg0 flagarg1", "flagarg2 flagarg3"}
				flag0, ok := flags[wantType[0]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[0])
					return
				}
				flag1, ok := flags[wantType[1]]
				if !ok {
					t.Errorf("flag with type %s not found", wantType[1])
					return
				}
				if strings.Compare(flag0.Type, wantType[0]) != 0 || strings.Compare(flag1.Type, wantType[1]) != 0 {
					t.Errorf("flag type is incorrect want %v, got %s", wantType, []string{flag0.Type, flag1.Type})
					return
				}
				if strings.Compare(flag0.Name, wantName[0]) != 0 || strings.Compare(flag1.Name, wantName[1]) != 0 {
					t.Errorf("flag name is incorrect want %v, got %v", wantName, []string{flag0.Name, flag1.Name})
					return
				}
				if len(args) != 2 || args[0] != wantArgs[0] || args[1] != wantArgs[1] {
					t.Errorf("argument is incorrect want = %v, got = %v", wantArgs, args)
				}
				if strings.Compare(flag0.Args, wantFlagArg[0]) != 0 || strings.Compare(flag1.Args, wantFlagArg[1]) != 0 {
					t.Errorf("flag args is incorrect want %v, got %v", wantFlagArg, []string{flag0.Args, flag1.Args})
					return
				}
			},
		},
	)
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		c       Cli
		args    args
		wantErr bool
	}{
		{
			name: "must run",
			c: cli,
			args: args {
				input: "do",
			},
			wantErr: false,
		},
		{
			name: "must return err 'command not found'",
			c: cli,
			args: args {
				input: "notdo",
			},
			wantErr: true,
		},
		{
			name: "cmd with arg",
			c: cli,
			args: args {
				input: "witharg arg",
			},
			wantErr: false,
		},
		{
			name: "cmd with two args",
			c: cli,
			args: args {
				input: "withtwoargs arg0 arg1",
			},
			wantErr: false,
		},
		{
			name: "cmd with multiword arg",
			c: cli,
			args: args {
				input: "withmultiwordarg \"arg0 arg1\"",
			},
			wantErr: false,
		},
		{
			name: "cmd with multiword args",
			c: cli,
			args: args {
				input: "withmultiwordargs \"arg0 arg1\" 'arg2 arg3'",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flag",
			c: cli,
			args: args {
				input: "withlongflag --flag",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flag",
			c: cli,
			args: args {
				input: "withlongargflag --flag=flagarg",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flags",
			c: cli,
			args: args {
				input: "withlongflags --flag0 --flag1",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flags with args",
			c: cli,
			args: args {
				input: "withlongflags --flag0=flagarg0 --flag1=flagarg1",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flags with multiword args",
			c: cli,
			args: args {
				input: "withlongflagswithmargs --flag0=\"flagarg0 flagarg1\" --flag1='flagarg2 flagarg3'",
			},
			wantErr: false,
		},
		{
			name: "cmd with short flag",
			c: cli,
			args: args {
				input: "withshortflag -f",
			},
			wantErr: false,
		},
		{
			name: "cmd with short flags",
			c: cli,
			args: args {
				input: "withshortflags -f -a",
			},
			wantErr: false,
		},
		{
			name: "cmd with short combined flags",
			c: cli,
			args: args {
				input: "withshortflags -fa",
			},
			wantErr: false,
		},
		{
			name: "cmd with logn flag and arg",
			c: cli,
			args: args {
				input: "withlongflagandarg --flag arg",
			},
			wantErr: false,
		},
		{
			name: "cmd with short flag and arg",
			c: cli,
			args: args {
				input: "withshortflagandarg -f arg",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flag and multiword args",
			c: cli,
			args: args {
				input: "withlongflagandmultiwordargs --flag \"arg0 arg1\" 'arg2 arg3'",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flag and multiword args",
			c: cli,
			args: args {
				input: "withshortflagandmultiwordargs -f \"arg0 arg1\" 'arg2 arg3'",
			},
			wantErr: false,
		},
		{
			name: "cmd with long flags with multiword args and multiword args",
			c: cli,
			args: args {
				input: "wlfwmaama --flag0=\"flagarg0 flagarg1\" --flag1='flagarg2 flagarg3' \"arg0 arg1\" 'arg2 arg3'",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.OneCmd(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("cli.OneCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
