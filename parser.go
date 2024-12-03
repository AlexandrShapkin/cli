package cli

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// CommandParser implements a method for parsing commands of the form <command> <flags> <args> into a ParsedCommand structure
type CommandParser interface {
	ParseCommand(input string) (*ParsedCommand, error) // ParseCommand parses a string representing a command of the form <command> <flags> <args>
}

type commandParser struct {
}

func NewCommandParser() CommandParser {
	return &commandParser{}
}

// ParseCommand parses a string representing a command of the form <command> <flags> <args> into a ParsedCommand structure
//
// Command examples:
// 	- cmd arg arg	- command with two splittted args
//	- cmd "arg arg" - command with one multiword flag
//	- cmd -a --flag1="value1" -bc --flag2='value2' -d 'arg arg' "arg arg"	- command with variable flags and two multiword args
func (cp *commandParser) ParseCommand(input string) (*ParsedCommand, error) {
	tokens, err := Tokenize(input)
	if err != nil {
		return nil, err
	}
	if len(tokens) == 0 {
		return nil, errors.New("empty input")
	}

	cmd := &ParsedCommand {
		Name: tokens[0],
		Flags: make(map[string]*ParsedCommandFlags),
	}

	i := 1
	for i < len(tokens) {
		token := tokens[i]
		if strings.HasPrefix(token, "--") {
			token = token[2:]
			parts := splitFlag(token)
			var arg string
			flag := parts[0]
			if len(parts) == 2 {
				arg = trimArg(parts[1])
			}
			cmd.Flags[flag] = &ParsedCommandFlags {
				Name: flag,
				Args: arg,
			}
		} else if strings.HasPrefix(tokens[i], "-") {
			token = token[1:]
			if len(token) == 1 {
				parts := splitFlag(token)
				var arg string
				flag := parts[0]
				if len(parts) == 2 {
					arg = trimArg(parts[1])
				}
				cmd.Flags[flag] = &ParsedCommandFlags {
					Name: flag,
					Args: arg,
				}
			} else {
				for _, f := range token {
					cmd.Flags[string(f)] = &ParsedCommandFlags{
						Name: string(f),
						Args: "",
					}
				}
			}
		} else {
			break
		}
		i++
	}

	cmd.Args = tokens[i:]

	return cmd, nil
}

func splitFlag(token string) []string {
	parts := strings.SplitN(token, "=", 2)
	return parts
}

func trimArg(arg string) string {
	if strings.HasPrefix(arg, "\"") && strings.HasSuffix(arg, "\"") {
		return strings.Trim(arg, "\"")
	} else if strings.HasPrefix(arg, "'") && strings.HasSuffix(arg, "'") {
		return strings.Trim(arg, "'")
	}
	return arg
}

// Tokenize split command input to string array
//
// Removes quotes where required
// Example:
//	- cmd --flag1="value1" "arg arg" -> []string{"cmd", "--flag1=value1", "arg arg"}
func Tokenize(input string) ([]string, error) {
	var tokens []string
	var currentToken strings.Builder
	inQuotes := false
	quoteRune := rune(0)

	for _, r := range input {
		switch {
		case r == '\'' || r == '"':
			if inQuotes {
				if r == quoteRune {
					inQuotes = false
					tokens = append(tokens, currentToken.String())
					currentToken.Reset()
				} else {
					currentToken.WriteRune(r)
				}
			} else {
				inQuotes = true
				quoteRune = r
			}
		case unicode.IsSpace(r) && !inQuotes:
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
		default:
			currentToken.WriteRune(r)
		}
	}

	if currentToken.Len() > 0 {
		if inQuotes {
			return nil, fmt.Errorf("unclosed quote at position %d", len(input))
		}
		tokens = append(tokens, currentToken.String())
	}

	return tokens, nil
}