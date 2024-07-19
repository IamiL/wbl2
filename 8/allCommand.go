package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

//Описание работы каждой из команд

const (
	CdCommand   = "cd"
	EchoCommand = "echo"
	KillCommand = "kill"
	PwdCommand  = "pwd"
	PsCommand   = "ps"
	QuitCommand = "quit"
)

// Command - общий интерфейс для команд
type Command interface {
	Execute(args ...string) ([]byte, error)
}

type commandCd struct{}

func (c *commandCd) Execute(args ...string) ([]byte, error) {
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return nil, err
	}

	dir, err = os.Getwd()
	if err != nil {
		return nil, err
	}
	return []byte("success cd " + dir), err
}

type commandPwd struct{}

func (c *commandPwd) Execute(args ...string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	return []byte(dir), err
}

type commandKill struct{}

func (c *commandKill) Execute(args ...string) ([]byte, error) {
	pId, err := strconv.Atoi(args[0])
	if err != nil {
		return nil, err
	}
	process, err := os.FindProcess(pId)
	if err != nil {
		return nil, err
	}
	err = process.Kill()
	if err != nil {
		return nil, err
	}
	return []byte("process be killed"), nil
}

type commandPs struct{}

func (c *commandPs) Execute(args ...string) ([]byte, error) {
	cmd := exec.Command("cmd.exe", "/c tasklist")
	return cmd.Output()
}

type commandEcho struct{}

func (c *commandEcho) Execute(args ...string) ([]byte, error) {
	prefix := strings.Split("/c echo", " ")
	args = append(prefix, args...)
	cmd := exec.Command("cmd.exe", args...)
	return cmd.Output()
}

type ShellUtil struct {
	command Command
	output  io.Writer
}

func (u *ShellUtil) SetCommand(command Command) {
	u.command = command
}

func (u *ShellUtil) ExecuteCommand(args ...string) {
	bytes, err := u.command.Execute(args...)
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
	_, err = fmt.Fprintf(u.output, "%s", string(bytes))
	if err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
}
func (u *ShellUtil) ExecuteCommands(commands []string) {
	for _, command := range commands {
		args := strings.Split(command, " ")

		com := args[0]
		if len(args) > 1 {
			args = args[1:]
		}

		switch com {
		case EchoCommand:
			cmd := &commandEcho{}
			u.SetCommand(cmd)
		case CdCommand:
			cmd := &commandCd{}
			u.SetCommand(cmd)
		case KillCommand:
			cmd := &commandKill{}
			u.SetCommand(cmd)
		case PwdCommand:
			cmd := &commandPwd{}
			u.SetCommand(cmd)
		case PsCommand:
			cmd := &commandPs{}
			u.SetCommand(cmd)
		case QuitCommand:
			_, err := fmt.Fprintln(u.output, "exit")
			if err != nil {
				fmt.Println("error: ", err.Error())
				return
			}
			os.Exit(1)
		default:
			fmt.Println("Такой команды нет")
			continue
		}
		u.ExecuteCommand(args...)
	}
}
