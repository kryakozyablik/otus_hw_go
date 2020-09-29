package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	if dir[len(dir)-1] != '/' {
		dir += "/"
	}

	env := make(Environment)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.Contains(file.Name(), "=") {
			continue
		}

		line, err := getLineFromFile(dir + file.Name())
		if err != nil {
			return nil, err
		}

		line = strings.ReplaceAll(line, "\x00", "\n")
		line = strings.TrimRight(line, "\t ")

		env[file.Name()] = line
	}

	return env, nil
}

func getLineFromFile(file string) (string, error) {
	f, err := os.Open(file)
	if err != nil {
		return "", err
	}

	fr := bufio.NewReader(f)
	line, err := fr.ReadString('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	err = f.Close()
	if err != nil {
		return "", err
	}

	return strings.TrimRight(line, "\n"), nil
}
