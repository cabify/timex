package timexmock

import "github.com/cabify/timex"

// Mocked runs the provided function with timex mocked, and then restores
// the default implementation
func Mocked(f func(*Implementation)) {
	mocked := &Implementation{}
	restore := timex.Override(mocked)
	defer restore()
	f(mocked)
}
