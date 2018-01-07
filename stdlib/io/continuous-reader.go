package main

import (
	"io"
	"log"
	"strings"
)

type continuousReader struct {
	target  string
	reader  io.Reader
	counter uint8
}

func (r continuousReader) newReader(t string) continuousReader {
	return continuousReader{
		target:  t,
		reader:  strings.NewReader(t),
		counter: 0,
	}
}

func (r continuousReader) Read(b []byte) (c int, err error) {
	if c, err = r.reader.Read(b); err != nil {
		log.Print(r.counter)
		if r.counter == 1 {
			nr := strings.NewReader("\nsurprise!\n")
			return nr.Read(b)
		}
		r.counter = 1
		return c, nil
	}
	return
}
