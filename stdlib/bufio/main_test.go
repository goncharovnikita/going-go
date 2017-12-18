package main_test

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestScanBytes(t *testing.T) {
	var (
		data = "test string"
		err  error
	)
	scanner := bufio.NewScanner(strings.NewReader(data))
	split := func(data []byte, atEof bool) (advance int, token []byte, e error) {
		advance, token, e = bufio.ScanBytes(data, atEof)
		return
	}

	scanner.Split(split)

	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
