package timex

import (
	"time"
)

// Timer is an interface similar to what time.Timer provides, but with C as a function returning the channel
// instead of accessing a plain struct property, so we can mock it.
// See time.Timer docs for more details
type Timer interface {
	// C returns the channel where the tick will be signaled, like Timer.C in time package
	C() <-chan time.Time

	// Stop stops the timer, see exact docs on time.Timer for information about channel draining when Stop returns false
	Stop() bool
}

//go:generate mockery -case underscore -outpkg timexmock -output timexmock -name Timer
type timer struct {
	*time.Timer
}

func (t timer) C() <-chan time.Time { return t.Timer.C }
