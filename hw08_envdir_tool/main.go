package main

import (
	"errors"
	"log"
	"os"
)

var (
	ErrNeedArguments = errors.New("please set 'env dir' and 'command'")
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal(ErrNeedArguments)
	}
	envDir := args[1]
	envList, err := ReadDir(envDir)
	if err != nil {
		log.Fatal(err)
	}

	returnCode := RunCmd(args[2:], envList)

	os.Exit(returnCode)
}
