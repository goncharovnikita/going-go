package main

import (
	"io"
	"os"
	"strings"
)

func main() {
	var reader continuousReader
	r := reader.newReader("some test string\n")
	io.Copy(os.Stdout, r)
}

// ReaderType type
type ReaderType struct {
	someData []byte
	reader   io.Reader
}

func (t ReaderType) Read(b []byte) (int, error) {
	r := strings.NewReader(string(t.someData))
	return r.Read(b)
}
