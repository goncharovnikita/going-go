package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No command passed")
	}
	cmd := os.Args[1]
	switch cmd {
	case "chdir":
		ChDir(os.Args[2])
		break
	}
}

// ChDir func
func ChDir(dirname string) {
	log.Fatal(os.Chdir(dirname))
}
