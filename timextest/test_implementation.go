package timextest

import (
	"sync"
	"time"

	"github.com/cabify/timex"
)

// Mock mocks/stubs the timex functions so they return constant known values
// Also mocks the timex.AfterFunc, returning the calls made to it through a channel.
// It mocks Sleep too: the delayed function will just unlock the caller of Sleep.
func Mock(now time.Time) *TestImplementation {
	impl := &TestImplementation{
		SleepCalls:     make(chan SleepCall),
		AfterCalls:     make(chan AfterCall),
		AfterFuncCalls: make(chan AfterFuncCall),
		NewTickerCalls: make(chan NewTickerCall),
		NewTimerCalls:  make(chan NewTimerCall),

		now: now,
	}

	impl.restore = timex.Override(impl)

	return impl
}

// Mocked mocks the current time and passes it to the provided function
// Afterwards, it restores the default implementation
func Mocked(now time.Time, f func(mocked *TestImplementation)) {
	mocked := Mock(now)
	defer mocked.TearDown()
	f(mocked)
}

// TestImplementation implements timex.Implementation for tests
type TestImplementation struct {
	SleepCalls     chan SleepCall
	AfterCalls     chan AfterCall
	AfterFuncCalls chan AfterFuncCall
	NewTickerCalls chan NewTickerCall
	NewTimerCalls  chan NewTimerCall

	sync.RWMutex
	now time.Time

	restore func()
}

// TearDown closes all the channels on the test implementation.
// Useful to terminate goroutines watching those channels.
func (ti *TestImplementation) TearDown() {
	close(ti.AfterFuncCalls)
	close(ti.AfterCalls)
	close(ti.NewTickerCalls)
	close(ti.NewTimerCalls)
	ti.RestoreDefaultImplementation()
}

// RestoreDefaultImplementation restores timex default implementation
func (ti *TestImplementation) RestoreDefaultImplementation() {
	ti.restore()
}

// SetNow updates Now() value to a newer one
func (ti *TestImplementation) SetNow(now time.Time) {
	ti.Lock()
	defer ti.Unlock()
	ti.now = now
}

// Now returns always the same now
func (ti *TestImplementation) Now() time.Time {
	ti.RLock()
	defer ti.RUnlock()
	return ti.now
}

// Since returns the duration elapsed since mocked `now`
func (ti *TestImplementation) Since(t time.Time) time.Duration {
	ti.RLock()
	defer ti.RUnlock()
	return ti.now.Sub(t)
}

// Until returns the duration until mocked `now`
func (ti *TestImplementation) Until(t time.Time) time.Duration {
	ti.RLock()
	defer ti.RUnlock()
	return t.Sub(ti.now)
}

// Sleep sleeps until WakeUp is called
func (ti *TestImplementation) Sleep(d time.Duration) {
	var wg sync.WaitGroup
	wg.Add(1)
	ti.SleepCalls <- SleepCall{Duration: d, WakeUp: wg.Done}
	wg.Wait()
}

// After returns a mocked timer
func (ti *TestImplementation) After(d time.Duration) <-chan time.Time {
	timer := newMockedTimer()
	ti.AfterCalls <- AfterCall{Mock: timer, Duration: d}
	return timer.C()
}

// AfterFunc allows the AfterFunc mocking by pushing calls into ti.AfterFuncCalls
func (ti *TestImplementation) AfterFunc(d time.Duration, f func()) timex.Timer {
	timer := newMockedFuncTimer(f)
	ti.AfterFuncCalls <- AfterFuncCall{Duration: d, Function: f, Mock: timer}
	return timer
}

// NewTicker returns a mocked ticker
func (ti *TestImplementation) NewTicker(d time.Duration) timex.Ticker {
	ticker := newMockedTicker()
	ti.NewTickerCalls <- NewTickerCall{Duration: d, Mock: ticker}
	return ticker
}

// NewTimer returns a mocked timer
func (ti *TestImplementation) NewTimer(d time.Duration) timex.Timer {
	timer := newMockedTimer()
	ti.NewTimerCalls <- NewTimerCall{Duration: d, Mock: timer}
	return timer
}

// AfterCall is a call to a mocked After
// It relies on a MockedTimer which will never be stopped
type AfterCall struct {
	Mock     *MockedTimer
	Duration time.Duration
}

// NewTickerCall is a call to a mocked NewTicker
type NewTickerCall struct {
	Mock     *MockedTicker
	Duration time.Duration
}

// NewTimerCall is a call to a mocked NewTimer
type NewTimerCall struct {
	Mock     *MockedTimer
	Duration time.Duration
}

// AfterFuncCall represents a call made to timex.AfterFunc
type AfterFuncCall struct {
	// Mock provides the underlying timer. Calling Trigger() on it will execute the function provided
	// in the calling goroutine
	Mock     *MockedTimer
	Duration time.Duration
	// Function is the function provided, probably useless
	Function func()
}

// SleepCall represents a call made to timex.Sleep
type SleepCall struct {
	// WakeUp allows calling goroutine to continue execution
	WakeUp   func()
	Duration time.Duration
}
