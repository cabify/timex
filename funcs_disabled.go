// +build timex_disable

package timex

import (
	"time"
)

// Now can be used as a replacement of time.Now()
func Now() time.Time { return time.Now() }

// Since can be used as a replacement of time.Since()
func Since(t time.Time) time.Duration { return time.Since(t) }

// Until can be used as a replacement of time.Until()
func Until(t time.Time) time.Duration { return time.Until(t) }

// Sleep can be used as a replacement of time.Sleep()
func Sleep(d time.Duration) { time.Sleep(d) }

// After can be used a replacement of time.After()
func After(d time.Duration) <-chan time.Time { return time.After(d) }

// AfterFunc can be used as a replacement of time.AfterFunc()
func AfterFunc(d time.Duration, f func()) Timer { return timer{time.AfterFunc(d, f)} }

// NewTicker creates a new ticker as a replacement of time.NewTicker
func NewTicker(d time.Duration) Ticker { return ticker{time.NewTicker(d)} }

// NewTimer creates a new timer as a replacement of time.NewTimer
func NewTimer(d time.Duration) Timer { return timer{time.NewTimer(d)} }
