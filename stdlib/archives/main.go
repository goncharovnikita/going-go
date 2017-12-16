package main

import (
	"archive/tar"
	"bytes"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		printUsage()
		return
	}
	c := getCommand(os.Args[1])
	switch c {
	case "tar":
		if err := Tarize(args[2]); err != nil {
			log.Fatalln(err)
		}
		break
	default:
		printUsage()
	}
}

// Tarize create tar file from existing
func Tarize(filename string) error {
	var err error
	var file *os.File
	var buf = new(bytes.Buffer)
	var tarizedName = GetTarizedName(filename)
	var archivedFile []byte
	// Open file
	if file, err = os.OpenFile(filename, os.O_RDONLY, 0644); err != nil {
		return err
	}
	// Read from file
	buf.ReadFrom(file)
	// Close file
	if err = file.Close(); err != nil {
		return err
	}

	// Compress to tar
	if archivedFile, err = WriteToArchive(filename, buf.Bytes()); err != nil {
		return err
	}

	// Open or create tar archive
	if file, err = os.OpenFile(tarizedName, os.O_RDWR|os.O_CREATE, 0644); err != nil {
		return err
	}

	// Write to archive
	if _, err = file.Write(archivedFile); err != nil {
		return err
	}

	// Close archive
	if err = file.Close(); err != nil {
		return err
	}

	return err
}

func isTar(filename string) bool {
	return strings.HasSuffix(filename, ".tar")
}

// WriteToArchive compress file to tar
func WriteToArchive(filename string, file []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	var err error

	if err = tw.WriteHeader(&tar.Header{
		Name: filename,
		Mode: 0600,
		Size: int64(len(file)),
	}); err != nil {
		return nil, err
	}

	if _, err = tw.Write(file); err != nil {
		return nil, err
	}

	if err = tw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
