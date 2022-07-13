package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTelnetClient(t *testing.T) {
	t.Run("success case", func(t *testing.T) {
		listener, err := net.Listen("tcp", "127.0.0.1:")
		require.NoError(t, err)
		defer func() { require.NoError(t, listener.Close()) }()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			in := &bytes.Buffer{}
			out := &bytes.Buffer{}

			timeout, err := time.ParseDuration("10s")
			require.NoError(t, err)

			client := NewTelnetClient(listener.Addr().String(), timeout, ioutil.NopCloser(in), out)
			require.NoError(t, client.Connect())
			defer func() { require.NoError(t, client.Close()) }()

			in.WriteString("hello\n")
			err = client.Send()
			require.NoError(t, err)

			err = client.Receive()
			require.NoError(t, err)
			require.Equal(t, "world\n", out.String())
		}()

		go func() {
			defer wg.Done()

			conn, err := listener.Accept()
			require.NoError(t, err)
			require.NotNil(t, conn)
			defer func() { require.NoError(t, conn.Close()) }()

			request := make([]byte, 1024)
			n, err := conn.Read(request)
			require.NoError(t, err)
			require.Equal(t, "hello\n", string(request)[:n])

			n, err = conn.Write([]byte("world\n"))
			require.NoError(t, err)
			require.NotEqual(t, 0, n)
		}()

		wg.Wait()
	})

	t.Run("invalid ip address", func(t *testing.T) {
		in := &bytes.Buffer{}
		out := &bytes.Buffer{}

		timeout, _ := time.ParseDuration("10s")

		client := NewTelnetClient("127.0.0.1:4343", timeout, ioutil.NopCloser(in), out)
		require.Error(t, client.Connect())
	})

	t.Run("empty message", func(t *testing.T) {
		l, err := net.Listen("tcp", "127.0.0.1:")
		require.NoError(t, err)
		defer func() { require.NoError(t, l.Close()) }()

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()

			in := new(bytes.Buffer)
			out := new(bytes.Buffer)

			timeout, _ := time.ParseDuration("10s")

			client := NewTelnetClient(l.Addr().String(), timeout, ioutil.NopCloser(in), out)
			require.NoError(t, client.Connect())
			require.NoError(t, client.Close())
		}()

		go func() {
			defer wg.Done()

			conn, err := l.Accept()
			require.NoError(t, err)
			require.NotNil(t, conn)
			defer func() { require.NoError(t, conn.Close()) }()

			request := make([]byte, 1024)
			_, err = conn.Read(request)
			require.Equal(t, fmt.Errorf("EOF"), err)
		}()

		wg.Wait()
	})
}
