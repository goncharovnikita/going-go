package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	var fileBuff []byte
	var err error
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	fileBuff, err = ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	if err := tw.WriteHeader(&tar.Header{
		Name: "guntar.tgz",
		Mode: 0600,
		Size: len(fileBuff),
	})
	fmt.Printf("Result: %s\n", string(fileBuff))
}
