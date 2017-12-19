package main_test

import "testing"
import "bytes"
import "compress/gzip"
import "io"
import "github.com/stretchr/testify/assert"
import compress "github.com/goncharovnikita/going-go/stdlib/compress"

func TestCompressGzip(t *testing.T) {
	var (
		testData = "some test string"
		result   []byte
		err      error
		buf      = new(bytes.Buffer)
		gr       *gzip.Reader
	)

	if result, err = compress.CompressGzip([]byte(testData)); err != nil {
		t.Fatal(err)
	}

	reader := bytes.NewReader(result)
	if gr, err = gzip.NewReader(reader); err != io.EOF && err != nil {
		t.Fatal(err)
	}

	if _, err = buf.ReadFrom(gr); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testData, string(buf.Bytes()))
}
