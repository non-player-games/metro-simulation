package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/non-player-games/metro-simulation/ticker"
)

var wg sync.WaitGroup

func init() {
	// TODO: create sample set of data
}

func main() {
	duration := 5 * time.Second
	// TODO: replace duration to 1*time.Minute
	ticker := ticker.NewTicker(duration, testTicker)
	ticker.Run()

	// to continue processing until quit
	wg.Add(1)
	wg.Wait()
}

func testTicker(t time.Time) error {
	fmt.Println("current tick", t)
	return nil
}
