package main

import (
	"fmt"
	"strings"
)

// available archives commands
var availableCommands = map[string]string{
	"tar":   "\tarchives tar [filename]\n\t\tCompress file to tar",
	"untar": "\tarchives untar [filename]\n\t\tDecompress tar to current directory",
}

// print usage
func printUsage() {
	fmt.Printf("\n\tUsage:\n")
	for _, v := range availableCommands {
		fmt.Println(v)
	}
}

// get command from args
func getCommand(cmd string) string {
	var result string
	for k := range availableCommands {
		if k == cmd {
			result = k
			break
		}
	}
	return result
}

// GetTarizedName convert filename to .tar
func GetTarizedName(s string) string {
	splitted := strings.Split(s, ".")

	if len(splitted) == 1 {
		return s + ".tar"
	}
	return strings.Join(splitted[:len(splitted)-1], ".") + ".tar"
}

// GetUntarizedName remove .tar suffix
func GetUntarizedName(s string) string {
	return strings.TrimSuffix(s, ".tar")
}
