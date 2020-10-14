package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]string

var ErrDirExists = errors.New("exists directory")
var ErrInvalidFile = errors.New("invalid file found")

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	env := make(Environment)

	for _, file := range files {
		if file.IsDir() {
			return nil, ErrDirExists
		}

		if strings.Contains(file.Name(), "=") {
			return nil, ErrInvalidFile
		}

		line, err := getLineFromFile(filepath.Join(dir, file.Name()))
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
