package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

type TelnetClient interface {
	Connect() error
	io.Closer
	Send() error
	Receive() error
}

type client struct {
	address string
	timeout time.Duration
	in      io.ReadCloser
	out     io.Writer
	con     net.Conn
}

func NewTelnetClient(address string, timeout time.Duration, in io.ReadCloser, out io.Writer) TelnetClient {
	return &client{
		address: address,
		timeout: timeout,
		in:      in,
		out:     out,
	}
}

func (c *client) Connect() error {
	if c.in == nil {
		return fmt.Errorf("in is invalid")
	}
	if c.out == nil {
		return fmt.Errorf("out is invalid")
	}

	con, err := net.DialTimeout("tcp", c.address, c.timeout)
	if err != nil {
		return err
	}

	c.con = con

	return nil
}

func (c *client) Close() error {
	return c.con.Close()
}

func (c *client) Send() error {
	if _, err := io.Copy(c.con, c.in); err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	return nil
}

func (c *client) Receive() error {
	if _, err := io.Copy(c.out, c.con); err != nil {
		return fmt.Errorf("failed to receive message: %w", err)
	}

	return nil
}
