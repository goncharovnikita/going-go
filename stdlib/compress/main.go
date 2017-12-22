package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {}

// CompressGzip compress bytes to gzip
func CompressGzip(filename string, file []byte) (result []byte, err error) {
	var buf = new(bytes.Buffer)
	var gw = gzip.NewWriter(buf)

	gw.Name = filename

	if _, err = gw.Write(file); err != nil {
		return nil, err
	}

	if err = gw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), err
}

// DecompressGzip decompress gicen data
func DecompressGzip(data []byte) (result []byte, err error) {
	var (
		reader = bytes.NewReader(data)
		buf    = new(bytes.Buffer)
		gr     *gzip.Reader
	)

	if gr, err = gzip.NewReader(reader); err != nil {
		return nil, err
	}

	fmt.Println(buf)

	if _, err = io.Copy(buf, gr); err != nil {
		return nil, err
	}

	if err = gr.Close(); err != nil {
		return nil, err
	}

	result = buf.Bytes()

	return result, err
}

// Gzipify compress given data to .gz file
func Gzipify(filename string, data []byte) (err error) {
	var compressed []byte
	var file *os.File
	var gzFilename = filename + ".gz"

	if compressed, err = CompressGzip(filename, data); err != nil {
		return err
	}

	if file, err = os.OpenFile(gzFilename, os.O_RDWR|os.O_CREATE, 0644); err != nil {
		return err
	}

	if _, err = file.Write(compressed); err != nil {
		return err
	}

	if err = file.Close(); err != nil {
		return err
	}

	return err
}

// Gunzip ungzip compressed file
// func Gunzip(filename string) (result []byte, err error) {
// 	var (
// 		buf  = new(bytes.Buffer)
// 		gr   *gzip.Reader
// 		file *os.File
// 	)

// 	if file, err = os.OpenFile(filename, os.O_RDONLY, 0644); err != nil {
// 		return
// 	}

// 	return
// }
