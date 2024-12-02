package cli

import "testing"

func Test_cli_OneCmd(t *testing.T) {
	cli := NewCli()
	cli.AddCmd(
		&Command{
			Use: "do",
			Run: func() {},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.OneCmd(tt.args.input); (err != nil) != tt.wantErr {
				t.Errorf("cli.OneCmd() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
