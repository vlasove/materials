package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/vlasove/materials/tasks_2/utils/shell/internal/app/shell"
)

const (
	// Version ...
	Version = "1.0.0 dev"
	// Author ...
	Author = "vlasove"
	// AppName ...
	AppName = "shell"
)

var (
	errUnapropriateFlagProvided = errors.New("shell: unkown flag")
	help                        bool
	version                     bool
)

func init() {
	flag.BoolVar(&help, "help", false, "Показать помощь и выйти.")
	flag.BoolVar(&version, "version", false, "Показать версию и немного информации и выйти.")
}

func usage() {
	log.Printf(`ПРОСТАЯ ОБОЛОЧКА
ИСПОЛЬЗОВАНИЕ: ./shell [COMMAND] <data>
ДОСТУПНЫЕ КОМАНДЫ:
  cd <data>   - изменить директорию. Если <data>
			  не передан - сбрасывается до домашнего католога
  echo <data> - эхо для входящей строки
  ps	- показать информацию о процессах. Работает
  				только на UNIX системах
  pwd	- показать текущую рабочую директорию
  kill <data> - убить процесс с именем <data>`)
	flag.PrintDefaults()
}

func showUsageAndExit(exitCode int) {
	usage()
	os.Exit(exitCode)
}

func showVersionAndExit(exitCode int) {
	log.Printf("%s . Version: '%s' Developed by: %s .",
		AppName,
		Version,
		Author,
	)
	os.Exit(exitCode)
}

func main() {
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	if help {
		showUsageAndExit(0)
	}

	if version {
		showVersionAndExit(0)
	}
	args := flag.Args()
	if len(args) < 1 {
		showUsageAndExit(1)
	}

	s := shell.New(args)
	switch args[0] {
	case "cd":
		cdExecutor := &shell.CDExecutor{}
		s.SetExecutor(cdExecutor)
	case "echo":
		echoExecutor := &shell.EchoExecutor{}
		s.SetExecutor(echoExecutor)
	case "ps":
		psExecutor := &shell.PSExecutor{}
		s.SetExecutor(psExecutor)
	case "pwd":
		pwdExecutor := &shell.PWDExecutor{}
		s.SetExecutor(pwdExecutor)
	case "kill":
		killExecutor := &shell.KillProcessExecutor{}
		s.SetExecutor(killExecutor)
	default:
		log.Fatal(errUnapropriateFlagProvided)
	}

	res, err := s.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}
