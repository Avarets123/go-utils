package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	rt := NewRateLimitter(5, time.Second*2)

	for i := range 25 {
		rt.Process(i)
	}

	time.Sleep(time.Second * 5)

}

type rateCounter struct {
	maxRate     int
	currentRate int
	d           time.Duration
	ticker      *time.Ticker
	mutex       sync.Mutex
	unblockCh   chan struct{}
}

func newCounter(maxRate int, d time.Duration) *rateCounter {
	rc := rateCounter{
		maxRate:     maxRate,
		currentRate: 0,
		d:           d,
		ticker:      time.NewTicker(d),
		mutex:       sync.Mutex{},
		unblockCh:   make(chan struct{}, 1),
	}

	rc.clearRateByInterval()
	return &rc
}

func (rc *rateCounter) clearRateByInterval() {

	go func() {

		for {
			<-rc.ticker.C
			fmt.Println("Rpc count: ", rc.currentRate)
			rc.restart()

			select {
			case rc.unblockCh <- struct{}{}:
			case <-rc.unblockCh:
			}

			fmt.Println("Ticker ticked")
		}

	}()

}

func (rc *rateCounter) increment() {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()
	if rc.maxRate == rc.currentRate {
		<-rc.unblockCh
	}

	rc.currentRate++
}

func (rc *rateCounter) restart() {
	rc.currentRate = 0

}

type RateLimitter struct {
	rateCounter *rateCounter
}

func NewRateLimitter(rpcCount int, d time.Duration) *RateLimitter {
	return &RateLimitter{
		rateCounter: newCounter(rpcCount, d),
	}
}

func (r *RateLimitter) Process(req int) {
	r.rateCounter.increment()
	fmt.Printf("Request: %d\n", req)
}
