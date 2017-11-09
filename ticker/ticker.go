package ticker

import (
	"log"
	"time"
)

var LogicalTime = 0

// TickFn defines what needs to be run
type TickFn func(t time.Time) error

// Ticker allows to define a ticker function that ticker every duration
type Ticker struct {
	Duration time.Duration
	TickerFn TickFn
	QuitChan chan bool
	Ticker   *time.Ticker
}

// NewTicker constructs a ticker
func NewTicker(duration time.Duration, tickerFn TickFn) Ticker {
	return Ticker{
		Duration: duration,
		TickerFn: tickerFn,
		QuitChan: make(chan bool),
		Ticker:   time.NewTicker(duration),
	}
}

// Run starts a goroutine for ticker to execute a function every interval
func (t Ticker) Run() {
	go func() {
		for ct := range t.Ticker.C {
            LogicalTime++
			if err := t.TickerFn(ct); err != nil {
				// TODO: think of a way to do better error handling e.g. fails
				// the whole loop after certain times and report to caller
				log.Println("Ticker has trouble calling TickFn", err)
			}
		}
		select {
		case <-t.QuitChan:
			t.Ticker.Stop()
			return
		}
	}()
}

// Stop will stop the ticker from the run function
func (t Ticker) Stop() {
	go func() {
		t.QuitChan <- true
	}()
}
