package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func MainUnix() {
	scan := bufio.NewScanner(os.Stdin)
	output := os.Stdout
	shell := &ShellUtil{output: output}
	for {
		fmt.Println("\n\ncommand: ")
		if scan.Scan() {
			line := scan.Text()
			cmds := strings.Split(line, " | ")
			shell.ExecuteCommands(cmds)
		}
	}
}
