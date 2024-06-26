package repl

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joegrn/jsh/pkg/cmds"
)

func Repl() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		input = strings.TrimSuffix(input, "\n")

		cmd, args := parseCmd(input)

		switch cmd {
		case "exit":
			cmds.Exit()
		case "echo":
			cmds.Echo(args)
		case "type":
			cmds.Type(args[0])
		case "pwd":
			cmds.Pwd()
		case "cd":
			cmds.Cd(args[0])
		default:
			cmds.NotFound(cmd, args)
		}

	}
}

func parseCmd(input string) (string, []string) {
	cmds := strings.Split(input, " ")
	cmd := cmds[0]

	args := []string{}
	if len(cmds) > 1 {
		args = cmds[1:]
	}

	return cmd, args
}
