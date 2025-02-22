package tcpproxy

import "os"

type StdinRead struct{}

type StdoutWrite struct{}

func (in *StdinRead) Read(b []byte) (int, error) {
	return os.Stdin.Read(b)
}

func (out *StdoutWrite) Write(b []byte) (int, error) {
	return os.Stdout.Write(b)
}
