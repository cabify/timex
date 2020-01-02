package timex

import "time"

// defaultImpl uses time package functions
type defaultImpl struct{}

// Now can be used as a replacement of time.Now()
func (defaultImpl) Now() time.Time {
	return time.Now()
}

// Since can be used as a replacement of time.Since()
func (defaultImpl) Since(t time.Time) time.Duration {
	return time.Since(t)
}

// Until can be used as a replacement of time.Until()
func (defaultImpl) Until(t time.Time) time.Duration {
	return time.Until(t)
}

// AfterFunc can be used as a replacement of time.AfterFunc()
func (defaultImpl) AfterFunc(d time.Duration, f func()) Timer {
	return timer{time.AfterFunc(d, f)}
}

// Sleep can be used as a replacement of time.Sleep()
func (defaultImpl) Sleep(d time.Duration) {
	time.Sleep(d)
}

// After can be used a replacement of time.After()
func (defaultImpl) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}

// NewTicker creates a new ticker as replacement of time.NewTicker()
func (defaultImpl) NewTicker(d time.Duration) Ticker {
	return ticker{time.NewTicker(d)}
}

// NewTimer creates a new timer as replacement of time.NewTimer()
func (defaultImpl) NewTimer(d time.Duration) Timer {
	return timer{time.NewTimer(d)}
}
