package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Echo(input []string) {
	text := strings.Join(input, " ")

	fmt.Fprintf(os.Stdout, "%s\n", text)
}

func Exit() {
	os.Exit(0)
}

func Type(cmd string) {
	switch cmd {
	case "exit", "echo", "type", "pwd", "cd":
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd)
	default:
		env := os.Getenv("PATH")
		paths := strings.Split(env, ":")
		for _, path := range paths {
			exec := path + "/" + cmd
			_, err := os.Stat(exec)
			if err != nil {
				continue
			}
			fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, exec)
			return
		}
		fmt.Fprintf(os.Stdout, "%s: not found\n", cmd)
	}
}

func NotFound(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmd)
	}
}

func Pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(os.Stdout, "%s\n", dir)
}

func Cd(path string) {
	if path == "~" {
		path = os.Getenv("HOME")
	}

	isAbsolutePath := path[0] == '/'

	if !isAbsolutePath {
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		path = dir + "/" + path
	}

	err := os.Chdir(path)
	if err != nil {
		fmt.Fprintf(os.Stdout, "cd: %s: No such file or directory\n", path)
	}
}
