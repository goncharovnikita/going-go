package runner_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/goncharovnikita/going-go/concurrency/runner"
)

func TestRunner(t *testing.T) {
	r := runner.New(time.Duration(200) * time.Millisecond)

	for i := 0; i < 10; i++ {
		r.Add(func(id int) {
			fmt.Printf("exetuting task #%d...\n", id)

			time.Sleep(time.Duration(100) * time.Millisecond)
		})
	}

	err := r.Start()

	if err != nil {
		t.Fatal(err)
	}

	t.Log("program finished\n")
}
