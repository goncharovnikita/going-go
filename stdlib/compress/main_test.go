package main_test

import "testing"
import "bytes"
import "compress/gzip"

import "github.com/stretchr/testify/assert"
import compress "github.com/goncharovnikita/going-go/stdlib/compress"
import "os"

func TestCompressGzip(t *testing.T) {
	var (
		testData = "some test string"
		filename = "testFilename"
		result   []byte
		err      error
		buf      = new(bytes.Buffer)
		gr       *gzip.Reader
	)

	if result, err = compress.CompressGzip(filename, []byte(testData)); err != nil {
		t.Fatal(err)
	}

	t.Log(result)

	reader := bytes.NewReader(result)
	if gr, err = gzip.NewReader(reader); err != nil {
		t.Fatal(err)
	}

	if err = gr.Close(); err != nil {
		t.Fatal(err)
	}

	if _, err = buf.ReadFrom(gr); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testData, string(buf.Bytes()))
}

func TestDecompressGzip(t *testing.T) {
	var (
		data   = []byte("some test data")
		buf    = new(bytes.Buffer)
		gw     = gzip.NewWriter(buf)
		result []byte
		err    error
	)

	if _, err = gw.Write(data); err != nil {
		t.Fatal(err)
	}

	if err = gw.Close(); err != nil {
		t.Fatal(err)
	}

	if result, err = compress.DecompressGzip(buf.Bytes()); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, data, result)
}

func TestGzipify(t *testing.T) {
	var (
		filename   = "filename"
		gzFilename = "filename.gz"
		data       = "some test data"
		buf        = new(bytes.Buffer)
		b          = new(bytes.Buffer)
		file       *os.File
		gr         *gzip.Reader
		err        error
	)

	if err = compress.Gzipify(filename, []byte(data)); err != nil {
		t.Fatal(err)
	}

	if file, err = os.OpenFile(gzFilename, os.O_RDONLY, 0644); err != nil {
		t.Fatal(err)
	}

	if _, err = buf.ReadFrom(file); err != nil {
		t.Fatal(err)
	}

	if err = file.Close(); err != nil {
		t.Fatal(err)
	}

	reader := bytes.NewReader(buf.Bytes())
	if gr, err = gzip.NewReader(reader); err != nil {
		t.Fatal(err)
	}

	if _, err = b.ReadFrom(gr); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, data, string(b.Bytes()))

	os.Remove(gzFilename)

}
