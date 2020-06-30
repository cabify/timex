package timex

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefault_Now(t *testing.T) {
	diff := time.Since(Default{}.Now())
	assert.True(t, diff < time.Second)
	assert.True(t, diff > 0)
}

func TestDefault_Since(t *testing.T) {
	diff := Default{}.Since(time.Now())
	assert.True(t, diff < time.Second)
	assert.True(t, diff > 0)
}

func TestDefault_Until(t *testing.T) {
	diff := Default{}.Until(time.Now())
	assert.True(t, diff < 0)
	assert.True(t, diff > -time.Second)
}

func TestDefault_AfterFunc(t *testing.T) {
	timeout := time.After(time.Second)
	ok := make(chan struct{})

	Default{}.AfterFunc(time.Millisecond, func() { close(ok) })

	select {
	case <-ok:
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for AfterFunc")
	}
}

func TestDefault_Sleep(t *testing.T) {
	timeout := time.After(time.Second)
	ok := make(chan struct{})

	go func() {
		Default{}.Sleep(time.Millisecond)
		close(ok)
	}()

	select {
	case <-ok:
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for Sleep")
	}
}

func TestDefault_After(t *testing.T) {
	timeout := time.After(time.Second)
	ok := Default{}.After(time.Millisecond)

	select {
	case <-ok:
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for After")
	}
}

func TestDefault_NewTicker(t *testing.T) {
	timeout := time.After(time.Second)
	ticker := Default{}.NewTicker(100 * time.Millisecond)

	select {
	case <-ticker.C():
	// ok
	case <-timeout:
		t.Errorf("Timeout waiting for Ticker")
	}

	ticker.Stop()

	ok := Default{}.After(200 * time.Millisecond)

	select {
	case <-ticker.C():
		t.Errorf("Should not tick again since it's stopped")
	case <-ok:
		// ok
	}
}

func TestDefault_NewTimer(t *testing.T) {
	t.Run("tick", func(t *testing.T) {
		timeout := time.After(time.Second)
		timer := Default{}.NewTimer(100 * time.Millisecond)

		select {
		case <-timer.C():
		// ok
		case <-timeout:
			t.Errorf("Timeout waiting for Mock")
		}
	})

	t.Run("stop", func(t *testing.T) {
		timer := Default{}.NewTimer(100 * time.Millisecond)
		timer.Stop()

		ok := After(200 * time.Millisecond)

		select {
		case <-timer.C():
			t.Errorf("Should not tick since it's stopped")
		case <-ok:
			// ok
		}
	})
}
