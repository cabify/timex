package timex

import "time"

// Default uses time package functions
type Default struct{}

// Now can be used as a replacement of time.Now()
func (Default) Now() time.Time {
	return time.Now()
}

// Since can be used as a replacement of time.Since()
func (Default) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Until can be used as a replacement of time.Until()
func (Default) Until(t time.Time) time.Duration {
	return time.Until(t)
}

// AfterFunc can be used as a replacement of time.AfterFunc()
func (Default) AfterFunc(d time.Duration, f func()) Timer {
	return timer{time.AfterFunc(d, f)}
}

// Sleep can be used as a replacement of time.Sleep()
func (Default) Sleep(d time.Duration) {
	time.Sleep(d)
}

// After can be used a replacement of time.After()
func (Default) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

// NewTicker creates a new ticker as replacement of time.NewTicker()
func (Default) NewTicker(d time.Duration) Ticker {
	return ticker{time.NewTicker(d)}
}

// NewTimer creates a new timer as replacement of time.NewTimer()
func (Default) NewTimer(d time.Duration) Timer {
	return timer{time.NewTimer(d)}
}
