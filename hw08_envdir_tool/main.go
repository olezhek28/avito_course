package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalln("two arguments required")
	}

	env, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	code := RunCmd(os.Args[2:], env)
	os.Exit(code)
}
