package os

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
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
	case "rr":
		start := time.Now()
		ReadRecursive(os.Args[2], 1)
		fmt.Printf("Read recursive taken %s\n", time.Since(start))
		break
	}
}

// ChDir func
func ChDir(dirname string) {
	log.Fatal(os.Chdir(dirname))
}

// ReadRecursive func
func ReadRecursive(filename string, depth int) {
	var (
		file  *os.File
		info  os.FileInfo
		infos []os.FileInfo
		err   error
	)

	if file, err = os.Open(filename); err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	if info, err = file.Stat(); err != nil {
		log.Fatal(err)
	}

	if isDir := info.IsDir(); isDir == true {
		fmt.Printf("%s %s/\n", strings.Repeat("-", depth), info.Name())
		if infos, err = file.Readdir(0); err != nil {
			log.Fatal(err)
		}
		for _, v := range infos {
			ReadRecursive(fmt.Sprintf("%s/%s", filename, v.Name()), depth+1)
		}
	} else {
		fmt.Printf("%s %s\n", strings.Repeat("-", depth), info.Name())
	}
}

// AppendToFile appends text to file
func AppendToFile(filename string, text string) (err error) {
	var (
		file *os.File
	)

	if file, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0755); err != nil {
		return
	}

	defer file.Close()

	if _, err = file.Write([]byte(text)); err != nil {
		return
	}

	return
}

// ReadDir func
// func ReadDir(dirname string) {
// 	var (
// 		infos []os.FileInfo
// 		err error
// 	)
// }
