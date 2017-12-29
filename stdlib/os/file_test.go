package main_test

import "testing"
import "os"

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
