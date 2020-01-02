// +build !timex_disable

package timex

import (
	"sync"
	"sync/atomic"
	"time"
)

// Now can be used as a replacement of time.Now()
func Now() time.Time { return impl.Load().(Implementation).Now() }

// Since can be used as a replacement of time.Since()
func Since(t time.Time) time.Duration { return impl.Load().(Implementation).Since(t) }

// Until can be used as a replacement of time.Until()
func Until(t time.Time) time.Duration { return impl.Load().(Implementation).Until(t) }

// Sleep can be used as a replacement of time.Sleep()
func Sleep(d time.Duration) { impl.Load().(Implementation).Sleep(d) }

// After can be used a replacement of time.After()
func After(d time.Duration) <-chan time.Time { return impl.Load().(Implementation).After(d) }

// AfterFunc can be used as a replacement of time.AfterFunc()
func AfterFunc(d time.Duration, f func()) Timer { return impl.Load().(Implementation).AfterFunc(d, f) }

// NewTicker creates a new ticker as a replacement of time.NewTicker
func NewTicker(d time.Duration) Ticker { return impl.Load().(Implementation).NewTicker(d) }

// NewTimer creates a new timer as a replacement of time.NewTimer
func NewTimer(d time.Duration) Timer { return impl.Load().(Implementation).NewTimer(d) }

// overridden is a lock we take when we override the timex implementation
// since golang packages can run tests in different packages concurrently,
// we want to make sure that there are no two implementations overriding
// concurrently
// https://medium.com/@xcoulon/how-to-avoid-parallel-execution-of-tests-in-golang-763d32d88eec
var overridden sync.Mutex

// Override replaces the global implementation used for timex
// The returned function should be called to restore the default implementation
func Override(implementation Implementation) func() {
	overridden.Lock()
	impl.Store(implValue{implementation})
	return func() {
		impl.Store(implValue{defaultImpl{}})
		overridden.Unlock()
	}
}

// impl stores the current implementation being used
// we don't use the RWMutex to read the impl itself because atomic.Value is faster in a _frequent read - unfrequent write_
// scenario according to the docs https://tip.golang.org/pkg/sync/atomic/#Value
var impl atomic.Value

type implValue struct{ Implementation }

func init() {
	impl.Store(implValue{defaultImpl{}})
}
