package main

import (
	"log"

	"github.com/asticode/go-astilectron"
)

func main() {
	var (
		a   *astilectron.Astilectron
		err error
	)
	if a, err = astilectron.New(astilectron.Options{
		AppName:            "helloworder",
		AppIconDefaultPath: "./assets/",
		AppIconDarwinPath:  "./assets/",
		BaseDirectoryPath:  "./",
	}); err != nil {
		log.Fatal(err)
	}
	defer a.Close()

	a.Start()
}
