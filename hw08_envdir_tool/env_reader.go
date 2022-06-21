package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var ErrInvalidFilename = errors.New("invalid filename")

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

type Environment map[string]EnvValue

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	fileInfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}

	env := make(Environment, len(fileInfos))

	for _, fInfo := range fileInfos {
		if fInfo.IsDir() {
			continue
		}

		if strings.Contains(fInfo.Name(), "=") {
			return nil, ErrInvalidFilename
		}

		var val string
		val, err = getValueFromFile(path.Join(dir, fInfo.Name()))
		if err != nil {
			return nil, err
		}

		env[fInfo.Name()] = EnvValue{
			Value:      val,
			NeedRemove: fInfo.Size() == 0,
		}
	}

	return env, nil
}

func getValueFromFile(fullName string) (string, error) {
	f, err := os.Open(fullName)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			log.Println("failed to close file: %w", err)
		}
	}()

	buf := bufio.NewReader(f)
	line, err := buf.ReadBytes('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	val := strings.ReplaceAll(string(line), "\x00", "\n")
	val = strings.TrimRight(val, " \t\n")

	return val, nil
}
