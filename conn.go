package oao

import (
	"bytes"
	"context"
	"io"
	"net"
)

var Prefix byte = '+'
var Suffix byte = '\n'
var SuffixB []byte = []byte{Suffix}
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
	w := c.And.GetWriter()
	if w == nil {
		return io.EOF
	}

	b2 := bytes.NewBuffer(nil)
	b2.WriteByte(Prefix)
	b2.Write(b)
	if !bytes.HasSuffix(b, SuffixB) {
		b2.WriteByte(Suffix)
	}
	_, err := w.Write(b2.Bytes())
	if err != nil {
		return err
	}
	return w.Flush()
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
