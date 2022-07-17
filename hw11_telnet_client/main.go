package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var timeoutStr string
	flag.StringVar(&timeoutStr, "timeout", "10s", "Duration of connections")
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("required arguments \"host\" and \"port\" not define")
	}

	address := net.JoinHostPort(flag.Arg(0), flag.Arg(1))
	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		log.Fatalf("timeout is invalid: %s", err)
	}

	telnetClient := NewTelnetClient(address, timeout, os.Stdin, os.Stdout)

	err = telnetClient.Connect()
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}
	defer telnetClient.Close()

	res := make(chan error)

	go func(res chan error) {
		res <- telnetClient.Receive()
	}(res)

	go func(res chan error) {
		res <- telnetClient.Send()
	}(res)

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, syscall.SIGINT)
	defer signal.Stop(sigint)

	select {
	case <-sigint:
	case <-res:
		close(sigint)
	}
}
