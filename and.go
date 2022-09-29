package oao

import (
	"bufio"
	"io"
)

var _ io.ReadWriteCloser = (*And)(nil)

type And struct {
	stdin  io.WriteCloser
	stdout io.ReadCloser
	stderr io.ReadCloser
}

func (a *And) Write(p []byte) (n int, err error) {
	if a.stdin == nil {
		return 0, io.EOF
	}
	return a.stdin.Write(p)
}

func (a *And) Close() (err error) {
	if a.stdin != nil {
		if err = a.stdin.Close(); err != nil {
			return
		}
	}
	if a.stderr != nil {
		if err = a.stderr.Close(); err != nil {
			return
		}
	}
	if a.stdout != nil {
		if err = a.stdout.Close(); err != nil {
			return
		}
	}
	return nil
}

func (a *And) Read(p []byte) (n int, err error) {
	if a.stdout != nil && a.stderr != nil {
		return io.MultiReader(a.stdout, a.stderr).Read(p)
	}
	if a.stdout == nil && a.stderr != nil {
		return a.stderr.Read(p)
	}
	if a.stdout != nil && a.stderr == nil {
		return a.stdout.Read(p)
	}
	return 0, io.EOF
}

func (a *And) GetReader() *bufio.Reader {
	return bufio.NewReader(a)
}
func (a *And) GetWriter() *bufio.Writer {
	return bufio.NewWriter(a)
}

func (a *And) GetNewReadWriter() *bufio.ReadWriter {
	return bufio.NewReadWriter(bufio.NewReader(a), bufio.NewWriter(a))
}
