package oao

import (
	"context"
	"io"
	"net"
)

var Prefix byte = '+'
var DefaultFmsg = func(_ []byte) {}

type Conn struct {
	*And
}

func NewConn(and *And) *Conn {
	return &Conn{And: and}
}

func (c *Conn) SentMsg(b []byte) error {
	if c.And == nil {
		return net.ErrClosed
	}
	_, err := c.And.Write(append([]byte{Prefix}, b...))
	return err
}

func (c *Conn) RecMsg(ctx context.Context, f func(b []byte)) error {
	if c.And == nil {
		return net.ErrClosed
	}
	if f == nil {
		f = DefaultFmsg
	}

	r := c.And.GetReader()
	if r == nil {
		return io.EOF
	}
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			l, _, err := r.ReadLine()
			if err != nil {
				return err
			}
			if len(l) < 2 || l[0] != Prefix {
				continue
			}
			f(l[1:])
		}
	}
}
