package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g := errgroup.Group{}

	g.Go(func() error {
		printLater("Hello\n", time.Millisecond*100)
		return nil
	})
	g.Go(func() error {
		// printLater("World\n", time.Millisecond*100)
		return fmt.Errorf("world failed")
	})
	g.Go(func() error {
		printLater(os.Getenv("USER")+"\n", time.Millisecond*100)
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "error found: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Print("All work completed as expected")
}

func printLater(msg string, duration time.Duration) {
	time.Sleep(duration)
	fmt.Printf(msg)
}
