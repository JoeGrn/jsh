package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
			Exit()
		case "echo":
			Echo(args)
		case "type":
			Type(args[0])
		case "pwd":
			Pwd()
		case "cd":
			Cd(args[0])
		default:
			NotFound(cmd, args)
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
