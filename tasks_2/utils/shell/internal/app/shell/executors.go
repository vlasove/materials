package shell

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/mitchellh/go-ps"
	"github.com/shirou/gopsutil/process"
)

var (
	errEmptyEcho         = errors.New("shell: echo should have some data")
	errDifferentPlatform = errors.New("shell: this command works only on UNIX system")
	errNeedProcess       = errors.New("shell: kill command needs name of process")
	errCanNotKillProcess = errors.New("shell: this process can not be killed")
)

type Executor interface {
	Execute(s *Shell) (string, error)
}

type CDExecutor struct{}

func (c *CDExecutor) Execute(s *Shell) (string, error) {
	// сброс до домашнего каталога
	if len(s.Args) == 1 {
		cUser, err := user.Current()
		if err != nil {
			return "", err
		}
		if err := os.Chdir(cUser.HomeDir); err != nil {
			return "", err
		}
	} else {
		// тут как минимум есть указанная директория
		if err := os.Chdir(s.Args[1]); err != nil {
			return "", err
		}

	}

	// отдаем путь до полученной директории
	p, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return "Current Dir: " + p, nil
}

type EchoExecutor struct{}

func (e *EchoExecutor) Execute(s *Shell) (string, error) {
	// если аргумент не передан
	if len(s.Args) == 1 {
		return "", errEmptyEcho
	}
	return s.Args[1], nil
}

type PSExecutor struct{}

func (p *PSExecutor) Execute(s *Shell) (string, error) {
	processes, err := ps.Processes()
	if err != nil {
		return "", err
	}
	var builder strings.Builder
	builder.WriteString("PID\t|\tCOMMAND\n")
	builder.WriteString("---------------\n")
	for _, proc := range processes {
		builder.WriteString(
			fmt.Sprintf("%v\t|\t%v\n", proc.Pid(), proc.Executable()),
		)
	}
	builder.WriteString("---------------\n")
	return builder.String(), nil
}

type PWDExecutor struct{}

func (p *PWDExecutor) Execute(s *Shell) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return "Current work dir :" + path, err
}

type KillProcessExecutor struct{}

func (k *KillProcessExecutor) Execute(s *Shell) (string, error) {
	// проверим, что есть , что убивать
	if len(s.Args) < 2 {
		return "", errNeedProcess
	}
	processes, err := process.Processes()
	if err != nil {
		return "", err
	}
	for _, process := range processes {
		name, err := process.Name()
		if err != nil {
			return "", err
		}
		if name == s.Args[1] {
			if err := process.Kill(); err != nil {
				return "", errCanNotKillProcess
			}
		}
	}
	return fmt.Sprintf("Process %v successfully killed", s.Args[1]), nil
}
