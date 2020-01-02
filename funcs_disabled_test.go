// +build timex_disable

package timex_test

import (
	"testing"
	"time"

	"github.com/cabify/timex"
	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	diff := time.Since(timex.Now())
	assert.True(t, diff < time.Second)
	assert.True(t, diff > 0)
}

func TestSince(t *testing.T) {
	diff := timex.Since(time.Now())
	assert.True(t, diff < time.Second)
	assert.True(t, diff > 0)
}

func TestUntil(t *testing.T) {
	diff := timex.Until(time.Now())
	assert.True(t, diff < 0)
	assert.True(t, diff > -time.Second)
}

func TestAfterFunc(t *testing.T) {
	timeout := time.After(time.Second)
	ok := make(chan struct{})

	timex.AfterFunc(time.Millisecond, func() { close(ok) })

	select {
	case <-ok:
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for AfterFunc")
	}
}

func TestSleep(t *testing.T) {
	timeout := time.After(time.Second)
	ok := make(chan struct{})

	go func() {
		timex.Sleep(time.Millisecond)
		close(ok)
	}()

	select {
	case <-ok:
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for Sleep")
	}
}

func TestAfter(t *testing.T) {
	timeout := time.After(time.Second)
	ok := timex.After(time.Millisecond)

	select {
	case <-ok:
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for After")
	}
}

func TestNewTicker(t *testing.T) {
	timeout := time.After(time.Second)
	ticker := timex.NewTicker(100 * time.Millisecond)

	select {
	case <-ticker.C():
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for Ticker")
	}

	ticker.Stop()

	ok := timex.After(200 * time.Millisecond)

	select {
	case <-ticker.C():
		t.Errorf("Should not tick again since it's stopped")
	case <-ok:
		// ok
	}
}

func TestNewTimer(t *testing.T) {
	t.Run("tick", func(t *testing.T) {
		timeout := time.After(time.Second)
		timer := timex.NewTimer(100 * time.Millisecond)

		select {
		case <-timer.C():
		// ok
		case <-timeout:
			t.Errorf("Timeout waiting for Mock")
		}
	})

	t.Run("stop", func(t *testing.T) {
		timer := timex.NewTimer(100 * time.Millisecond)
		timer.Stop()

		ok := timex.After(200 * time.Millisecond)

		select {
		case <-timer.C():
			t.Errorf("Should not tick since it's stopped")
		case <-ok:
			// ok
		}
	})
}
