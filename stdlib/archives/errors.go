package main

import (
	"fmt"
)

type fileIsNotTarError struct {
	Filename string
}

func (f fileIsNotTarError) Error() string {
	return fmt.Sprintf("file %s is not tar", f.Filename)
}
