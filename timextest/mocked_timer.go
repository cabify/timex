package timextest

import (
	"sync"
	"time"
)

// MockedTimer implements a timex.Timer and provides functions to control its behavior
type MockedTimer struct {
	c chan time.Time
	f func()

	stopped chan struct{}

	stopMutex sync.Mutex
	stopValue bool
}

func newMockedTimer() *MockedTimer {
	timer := &MockedTimer{
		c:       make(chan time.Time),
		stopped: make(chan struct{}),
	}
	timer.stopMutex.Lock()
	return timer
}

func newMockedFuncTimer(f func()) *MockedTimer {
	timer := &MockedTimer{
		f:       f,
		stopped: make(chan struct{}),
	}
	timer.stopMutex.Lock()
	return timer
}

// C implements timex.Timer
func (mt *MockedTimer) C() <-chan time.Time {
	return mt.c
}

// Stop implements the timex.Timer, stop will not return until the value is set by StopValue()
func (mt *MockedTimer) Stop() bool {
	close(mt.stopped)
	mt.stopMutex.Lock()
	defer mt.stopMutex.Unlock()
	return mt.stopValue
}

// StopValue sets the value to be returned by Stop(), it has to be called to allow
// that method to return
// It can be only called once
func (mt *MockedTimer) StopValue(v bool) {
	mt.stopValue = v
	mt.stopMutex.Unlock()
}

// StoppedChan provides the channel that is closed when Stop() is called
func (mt *MockedTimer) StoppedChan() chan struct{} {
	return mt.stopped
}

// Trigger for a usual Timer will send the provided time through timer.C(),
// If this was a AfterFunc call, then trigger will ignore the time provided and will
// call the scheduled function in the caller's goroutine
func (mt *MockedTimer) Trigger(t time.Time) {
	select {
	case <-mt.stopped:
		panic("trying to tick on a stopped timer, does your test have a race condition?")
	default:
		if mt.c == nil {
			mt.f()
		} else {
			mt.c <- t
		}
	}
}
