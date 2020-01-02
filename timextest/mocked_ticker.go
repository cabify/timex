package timextest

import (
	"time"
)

// MockedTicker implements a timex.Ticker and provides functions to control its behavior
type MockedTicker struct {
	c       chan time.Time
	stopped chan struct{}
}

func newMockedTicker() *MockedTicker {
	return &MockedTicker{
		c:       make(chan time.Time),
		stopped: make(chan struct{}),
	}
}

// C implements timex.Ticker
func (mt *MockedTicker) C() <-chan time.Time {
	return mt.c
}

// Stop implements timex.Ticker, and it will close the channel provided by StoppedChan
func (mt *MockedTicker) Stop() {
	close(mt.stopped)
}

// StoppedChan will be closed once Stop is called
func (mt *MockedTicker) StoppedChan() chan struct{} {
	return mt.stopped
}

// Tick will send the provided time through ticker.C()
func (mt *MockedTicker) Tick(t time.Time) {
	select {
	case <-mt.stopped:
		panic("trying to tick on a stopped ticker, does your test have a race condition?")
	default:
		mt.c <- t
	}
}
