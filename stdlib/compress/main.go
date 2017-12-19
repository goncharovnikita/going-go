package main

import "bytes"
import "compress/gzip"

func main() {}

// CompressGzip compress bytes to gzip
func CompressGzip(file []byte) (result []byte, err error) {
	var buf = new(bytes.Buffer)
	var gw = gzip.NewWriter(buf)

	if _, err = gw.Write(file); err != nil {
		return nil, err
	}

	if err = gw.Close(); err != nil {
		return nil, err
	}

	return
}
