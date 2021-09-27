package command

import (
	"fmt"
	"strconv"
	"strings"
)

type MwArguments struct {
}

func (mw *MwArguments) Exec(ctx *Context, cmd Commander) (next bool, err error) {
	defer func() {
		if !next && err != nil {
			errMessage := err.Error()
			errMessage += fmt.Sprintf("\nUsage of [%s]:\n", strings.Join(cmd.Invokes(), ", "))
			errMessage += cmd.Usage()
			err = fmt.Errorf(errMessage)
		}
	}()

	err = cmd.NArgs(ctx.inputArgs)
	if err != nil{
		return
	}

	flagMap, argMap, err := parseArgs(ctx.inputArgs, cmd)
	if err != nil {
		return
	}

	ctx.FlagMap = flagMap
	ctx.Args = argMap

	next = true
	return
}

func parseArgs(args []string, cmd Commander) (flagMap map[string]string, nonFlag []string, err error) {
	flag := ""
	flagMap = map[string]string{}

	for _, arg := range args {
		if isFlag(arg) {
			if !isValidFlag(cmd, arg) {
				err = fmt.Errorf("%s is not a valid flag", arg)
				return
			}
			flag = arg[1:]
			flagMap[flag] = ""
		} else {
			if flag == "" {
				if !isValidArg(cmd, arg) {
					err = fmt.Errorf("%s is not a valid arg", arg)
					return
				}
				nonFlag = append(nonFlag, arg)
			} else {
				if !isValidArgOfFlag(cmd, flag, arg) {
					err = fmt.Errorf("%s is not a valid arg", arg)
					return
				}
				flagMap[flag] = arg
				flag = ""
			}
		}
	}

	return
}

func isFlag(arg string) bool {
	return (len(arg) >= 3 && arg[0] == '-' && arg[1] == '-') ||
		(len(arg) >= 2 && arg[0] == '-' && arg[1] != '-')
}
func isValidFlag(cmd Commander, arg string) bool {
	flagMap := cmd.ValidFlags()
	arg = arg[1:]

	if _, ok := flagMap[arg]; ok {
		return true
	}

	return false
}
func isValidArg(cmd Commander, arg string) bool {
	validArgs := cmd.ValidArgs()
	if validArgs == nil {
		return true
	}

	for _, validArg := range validArgs {
		if arg == validArg {
			return true
		}
	}
	return false
}
func isValidArgOfFlag(cmd Commander, flag, arg string) bool {
	flagMap := cmd.ValidFlags()
	validArgs := flagMap[flag]
	for _, validArg := range validArgs {
		if arg == validArg {
			return true
		}
		if strings.HasPrefix(validArg, "@int") {
			split := strings.Split(validArg, " ")
			min, _ := strconv.Atoi(split[1])
			max, _ := strconv.Atoi(split[2])
			if i, err := strconv.Atoi(arg); err == nil && i >= min && i <= max {
				return true
			}
		}
	}
	return false
}

func ExactArgs(n int) PositionalArgs {
	return func(args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d args, received %d", n, len(args))
		}
		return nil
	}
}
func MinimumNArgs(n int) PositionalArgs {
	return func(args []string) error {
		if len(args) < n {
			return fmt.Errorf("requires at least %d arg(s), only received %d", n, len(args))
		}
		return nil
	}
}
func MaximumNArgs(n int) PositionalArgs {
	return func(args []string) error {
		if len(args) > n {
			return fmt.Errorf("accepts at most %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}