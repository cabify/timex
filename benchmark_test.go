package timex_test

import (
	"testing"
	"time"

	"github.com/cabify/timex"
)

var dontOptimizeMePlease time.Time

func BenchmarkTimeNow(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		dontOptimizeMePlease = time.Now()
	}
}

func BenchmarkTimexNow(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		dontOptimizeMePlease = timex.Now()
	}
}
