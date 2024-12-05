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
			name: "cmd with short flag",
			c: cli,
			args: args {
				input: "withshortflag -f",
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.OneCmd(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("cli.OneCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
