package timex

import (
	"time"
)

// Ticker is an interface similar to what time.Ticker provides, but with C as a function returning the channel
// instead of accessing a plain struct property, so we can mock it.
// See time.Ticker docs for more details
type Ticker interface {
	// C returns the channel where each tick will be signaled, like Ticker.C in time package
	C() <-chan time.Time

	// Stop stops the ticker, no more ticks will be received in `C()`
	Stop()
}

//go:generate mockery -case underscore -outpkg timexmock -output timexmock -name Ticker
type ticker struct {
	*time.Ticker
}

func (t ticker) C() <-chan time.Time { return t.Ticker.C }
