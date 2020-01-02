package timex

import (
	"time"
)

// Implementation defines the methods we delegate
type Implementation interface {
	Now() time.Time
	Since(t time.Time) time.Duration
	Until(t time.Time) time.Duration

	Sleep(d time.Duration)

	After(d time.Duration) <-chan time.Time
	AfterFunc(d time.Duration, f func()) Timer

	NewTicker(d time.Duration) Ticker
	NewTimer(d time.Duration) Timer
}

//go:generate mockery -case underscore -outpkg timexmock -output timexmock -name Implementation
var _ Implementation = defaultImpl{}
