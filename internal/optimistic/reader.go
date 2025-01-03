package optimistic

import (
	"bufio"
	"bytes"
	"github.com/aivarasbaranauskas/aoc/internal/_a"
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
	return _a.E(r.Reader.ReadString(delim))
}

func (r *Reader) ReadBytes(delim byte) []byte {
	return _a.E(r.Reader.ReadBytes(delim))
}

func (r *Reader) ReadStringLine() string {
	return strings.Trim(r.ReadString('\n'), "\n")
}

func (r *Reader) ReadBytesLine() []byte {
	return bytes.Trim(r.ReadBytes('\n'), "\n")
}
