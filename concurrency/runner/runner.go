// Package runner implements concurrency pattern
package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner struct
type Runner struct {
	timeout   <-chan time.Time
	interrupt chan os.Signal
	complete  chan error
	tasks     []func(int)
}

// ErrTimeout signals that task takes too long
var ErrTimeout = errors.New("recieved timeout")

// ErrInterrupt signals that interrupt signal catched
var ErrInterrupt = errors.New("recieved interrupt signal")

// New factory for runners
func New(dur time.Duration) *Runner {
	return &Runner{
		timeout:   time.After(dur),
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
	}
}

// Add attaches tasks to runner
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start starts runner
func (r *Runner) Start() (err error) {
	// catch interrupt signals
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	case err = <-r.complete:
		return nil
	case <-r.timeout:
		return ErrTimeout
	}
}

// run exetutes tasks
func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		task(id)
	}

	return nil
}

// gotInterrupt verifies if the interrupt signal has been issued.
func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}
