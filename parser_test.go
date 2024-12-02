package cli

import (
	"reflect"
	"testing"
)

func Test_tokenize(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "without flags and quotes",
			args: args{
				input: "cmd arg arg",
			},
			want:    []string{"cmd", "arg", "arg"},
			wantErr: false,
		},
		{
			name: "without flags double quotes",
			args: args{
				input: "cmd \"arg arg\"",
			},
			want:    []string{"cmd", "arg arg"},
			wantErr: false,
		},
		{
			name: "without flags single quotes",
			args: args{
				input: "cmd 'arg arg'",
			},
			want:    []string{"cmd", "arg arg"},
			wantErr: false,
		},
		{
			name: "without flags single quotes and double quotes",
			args: args{
				input: "cmd 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "arg arg", "arg arg"},
			wantErr: false,
		},
		{
			name: "with short flags single quotes and double quotes",
			args: args{
				input: "cmd -a -bc 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "-a", "-bc", "arg arg", "arg arg"},
			wantErr: false,
		},
		{
			name: "with long flags single quotes and double quotes",
			args: args{
				input: "cmd --flag1 --flag2 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "--flag1", "--flag2", "arg arg", "arg arg"},
			wantErr: false,
		},
		{
			name: "with mixed flags single quotes and double quotes",
			args: args{
				input: "cmd -a --flag1 -bc --flag2 -d 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "-a", "--flag1", "-bc", "--flag2", "-d", "arg arg", "arg arg"},
			wantErr: false,
		},
		{
			name: "with mixed flags and values single quotes and double quotes",
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "-a", "--flag1=value1", "-bc", "--flag2=value2", "-d", "arg arg", "arg arg"},
			wantErr: false,
		},
		{
			name: "unclosed double quote in flag",
			args: args{
				input: "cmd -a --flag1=\"value1 -bc --flag2='value2' -d 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "-a", "--flag1=value1", "-bc", "--flag2=value2", "-d", "arg arg", "arg arg"},
			wantErr: true,
		},
		{
			name: "unclosed single quote in flag",
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2 -d 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "-a", "--flag1=value1", "-bc", "--flag2=value2", "-d", "arg arg", "arg arg"},
			wantErr: true,
		},
		{
			name: "unclosed single and double quote in flag",
			args: args{
				input: "cmd -a --flag1=\"value1 -bc --flag2='value2 -d 'arg arg' \"arg arg\"",
			},
			want:    []string{"cmd", "-a", "--flag1=value1", "-bc", "--flag2=value2", "-d", "arg arg", "arg arg"},
			wantErr: true,
		},
		{
			name: "unclosed single quote in value",
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg \"arg arg\"",
			},
			want:    []string{"cmd", "-a", "--flag1=value1", "-bc", "--flag2=value2", "-d", "arg arg", "arg arg"},
			wantErr: true,
		},
		{
			name: "unclosed double quote in value",
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg' \"arg arg",
			},
			want:    []string{"cmd", "-a", "--flag1=value1", "-bc", "--flag2=value2", "-d", "arg arg", "arg arg"},
			wantErr: true,
		},
		{
			name: "unclosed single and double quote in value",
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg \"arg arg",
			},
			want:    []string{"cmd", "-a", "--flag1=value1", "-bc", "--flag2=value2", "-d", "arg arg", "arg arg"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Tokenize(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_commandParser_ParseCommand(t *testing.T) {
	parser := NewCommandParser()
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		cp      CommandParser
		args    args
		want    *ParsedCommand
		wantErr bool
	}{
		{
			name: "without flags and quotes",
			cp:   parser,
			args: args{
				input: "cmd arg arg",
			},
			want: &ParsedCommand{
				Name:  "cmd",
				Flags: make(map[string]*ParsedCommandFlags),
				Args:  []string{"arg", "arg"},
			},
			wantErr: false,
		},
		{
			name: "without flags double quotes",
			cp:   parser,
			args: args{
				input: "cmd \"arg arg\"",
			},
			want: &ParsedCommand{
				Name:  "cmd",
				Flags: make(map[string]*ParsedCommandFlags),
				Args:  []string{"arg arg"},
			},
			wantErr: false,
		},
		{
			name: "without flags single quotes",
			cp:   parser,
			args: args{
				input: "cmd 'arg arg'",
			},
			want: &ParsedCommand{
				Name:  "cmd",
				Flags: make(map[string]*ParsedCommandFlags),
				Args:  []string{"arg arg"},
			},
			wantErr: false,
		},
		{
			name: "without flags single quotes and double quotes",
			cp:   parser,
			args: args{
				input: "cmd 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name:  "cmd",
				Flags: make(map[string]*ParsedCommandFlags),
				Args:  []string{"arg arg", "arg arg"},
			},
			wantErr: false,
		},
		{
			name: "with short flags single quotes and double quotes",
			cp:   parser,
			args: args{
				input: "cmd -a -bc 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: false,
		},
		{
			name: "with long flags single quotes and double quotes",
			cp:   parser,
			args: args{
				input: "cmd --flag1 --flag2 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"flag1": {
						Name: "flag1",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: false,
		},
		{
			name: "with mixed flags single quotes and double quotes",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1 -bc --flag2 -d 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: false,
		},
		{
			name: "with mixed flags and values single quotes and double quotes",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "value1",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "value2",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: false,
		},
		{
			name: "unclosed double quote in flag",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1=\"value1 -bc --flag2='value2' -d 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "value1",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "value2",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: true,
		},
		{
			name: "unclosed single quote in flag",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2 -d 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "value1",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "value2",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: true,
		},
		{
			name: "unclosed single and double quote in flag",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1=\"value1 -bc --flag2='value2 -d 'arg arg' \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "value1",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "value2",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: true,
		},
		{
			name: "unclosed single quote in value",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg \"arg arg\"",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "value1",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "value2",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: true,
		},
		{
			name: "unclosed double quote in value",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg' \"arg arg",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "value1",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "value2",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: true,
		},
		{
			name: "unclosed single and double quote in value",
			cp:   parser,
			args: args{
				input: "cmd -a --flag1=\"value1\" -bc --flag2='value2' -d 'arg arg \"arg arg",
			},
			want: &ParsedCommand{
				Name: "cmd",
				Flags: map[string]*ParsedCommandFlags{
					"a": {
						Name: "a",
						Args: "",
					},
					"flag1": {
						Name: "flag1",
						Args: "value1",
					},
					"b": {
						Name: "b",
						Args: "",
					},
					"c": {
						Name: "c",
						Args: "",
					},
					"flag2": {
						Name: "flag2",
						Args: "value2",
					},
					"d": {
						Name: "d",
						Args: "",
					},
				},
				Args: []string{"arg arg", "arg arg"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.cp.ParseCommand(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("commandParser.ParseCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) && !tt.wantErr {
				t.Errorf("commandParser.ParseCommand() = %v, want %v", got, tt.want)
			}
		})
	}
}
