package optimistic

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type Reader struct {
	*bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewReader(r)}
}

func (r *Reader) ReadString(delim byte) string {
	s, err := r.Reader.ReadString(delim)
	if err != nil {
		panic(err)
	}
	return s
}

func (r *Reader) ReadBytes(delim byte) []byte {
	s, err := r.Reader.ReadBytes(delim)
	if err != nil {
		panic(err)
	}
	return s
}

func (r *Reader) ReadStringLine() string {
	return strings.Trim(r.ReadString('\n'), "\n")
}

func (r *Reader) ReadBytesLine() []byte {
	return bytes.Trim(r.ReadBytes('\n'), "\n")
}
