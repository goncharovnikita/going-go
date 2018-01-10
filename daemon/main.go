package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

var name = flag.String("name", "world", "name to be printing")

func init() {
	flag.Parse()
}

func main() {
	log.Printf("Starting service for %s\n", *name)

	signals := make(chan os.Signal, 1)

	signal.Notify(signals)

	go func() {
		s := <-signals
		log.Printf("Recieved signal %s\n", s)
		cleanup()
		os.Exit(1)
	}()

	for {
		log.Printf("Hello %s\n", *name)

		nsecs := rand.Intn(3000)
		log.Printf("Sleep for %dms before next loop\n", nsecs)
		time.Sleep(time.Millisecond * time.Duration(nsecs))
	}
}

func cleanup() {
	log.Println("Cleaning...")
}
