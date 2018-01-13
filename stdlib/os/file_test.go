package os_test

import (
	"io/ioutil"
	"os"
	"testing"

	gos "github.com/goncharovnikita/going-go/stdlib/os"
	"github.com/stretchr/testify/assert"
)

func TestFileOpen(t *testing.T) {
	var (
		err      error
		file     *os.File
		filename = "test"
	)

	if file, err = os.Open(filename); err != nil {
		t.Fatal(err)
	}

	defer file.Close()

}

func TestAppendToFile(t *testing.T) {
	var (
		file       *os.File
		err        error
		actualText string
		body       []byte
	)

	testText := "some test text"
	appendText := "appended text"
	expectedtext := "some test textappended text"
	filename := "testfile"

	if file, err = os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0755); err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	if _, err = file.Write([]byte(testText)); err != nil {
		t.Fatal(err)
	}

	if err = gos.AppendToFile(filename, appendText); err != nil {
		t.Fatal(err)
	}

	if file, err = os.OpenFile(filename, os.O_RDONLY, 0755); err != nil {
		t.Fatal(err)
	}

	defer file.Close()
	defer os.Remove(filename)

	if body, err = ioutil.ReadAll(file); err != nil {
		t.Fatal(err)
	}

	actualText = string(body)

	assert.Equal(t, expectedtext, actualText)
}
