package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	t, e := time.Parse("2006-01-02", "2018-01-14")
	if e != nil {
		log.Fatal(e)
	}

	now := time.Now()

	eq := compareDates(&t, &now)

	fmt.Print(eq)
}

func compareDates(v1 *time.Time, v2 *time.Time) bool {
	return v1.Year() == v2.Year() && v1.Month() == v2.Month() && v1.Day() == v2.Day()
}
