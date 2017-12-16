package main_test

import (
	"archive/tar"
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	archives "github.com/goncharovnikita/going-go/stdlib/archives"
)

func TestWriteToArchive(t *testing.T) {
	var filename = "test.txt"
	var file = []byte("Here some test text")
	var err error
	var result []byte

	if result, err = archives.WriteToArchive(filename, file); err != nil {
		t.Fatal(err)
	}

	var r = bytes.NewReader(result)
	var tr = tar.NewReader(r)
	var buf = new(bytes.Buffer)
	for {
		_, err = tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		buf.ReadFrom(tr)
	}

	assert.Equal(t, file, buf.Bytes())
}

func TestGetTarizedName(t *testing.T) {
	var cases = map[string]string{
		"test1.txt":        "test1.tar",
		"test2.nested.txt": "test2.nested.tar",
		".txt":             ".tar",
		"test4":            "test4.tar",
	}

	for k, v := range cases {
		assert.Equal(t, v, archives.GetTarizedName(k))
	}
}

func TestGetUntarizedName(t *testing.T) {
	var cases = map[string]string{
		"test1.tar":        "test1",
		"test2.nested.tar": "test2.nested",
		".tar":             "",
		"test4":            "test4",
	}

	for k, v := range cases {
		assert.Equal(t, v, archives.GetUntarizedName(k))
	}
}

func TestTarize(t *testing.T) {
	var testText = "some test text"
	var filename = "test.txt"
	var tarizedFilename = "test.tar"
	var err error
	var file *os.File

	os.Remove(filename)
	os.Remove(tarizedFilename)

	if file, err = os.Create(filename); err != nil {
		t.Fatal(err)
	}

	if _, err = file.Write([]byte(testText)); err != nil {
		t.Fatal(err)
	}

	if err = file.Close(); err != nil {
		t.Fatal(err)
	}

	if err = archives.Tarize(filename); err != nil {
		t.Fatal(err)
	}

	if file, err = os.OpenFile(tarizedFilename, os.O_RDONLY, 0600); err != nil {
		t.Fatal(err)
	}

	var buf = new(bytes.Buffer)
	var tr = tar.NewReader(file)
	for {
		_, err = tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if _, err = buf.ReadFrom(tr); err != nil {
			t.Fatal(err)
		}
	}

	if err = file.Close(); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, testText, string(buf.Bytes()))

	os.Remove(filename)
	os.Remove(tarizedFilename)

}
