package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var ErrInvalidFilename = errors.New("invalid filename")

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

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

		if fInfo.Size() == 0 {
			delete(env, fInfo.Name())
			continue
		}

		var val string
		val, err = getValueFromFile(path.Join(dir, fInfo.Name()))
		if err != nil {
			return nil, err
		}

		env[fInfo.Name()] = EnvValue{
			Value: val,
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
			fmt.Errorf("failed to close file: %w", err)
		}
	}()

	buf := bufio.NewReader(f)
	line, err := buf.ReadBytes('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	val := bytes.ReplaceAll(line, []byte{0}, []byte{'\n'})
	valStr := strings.TrimRight(string(val), " \t\n")

	return valStr, nil
}
