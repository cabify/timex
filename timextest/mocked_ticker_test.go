package timextest_test

import (
	"testing"
	"time"

	"github.com/cabify/timex"
	"github.com/cabify/timex/timextest"
	"github.com/stretchr/testify/assert"
)

func TestMockedTicker_StoppedChan(t *testing.T) {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		go func() {
			ticker := timex.NewTicker(time.Second)
			ticker.Stop()
		}()
		mockedTicker := <-mockedtimex.NewTickerCalls

		select {
		case <-mockedTicker.Mock.StoppedChan():
		case <-time.After(time.Second):
			t.Errorf("Stop should have been called")
		}
	})
}

func TestMockedTicker_Tick(t *testing.T) {
	t.Run("panics on stopped ticker", func(t *testing.T) {
		timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
			go func() {
				ticker := timex.NewTicker(time.Second)
				ticker.Stop()
			}()
			mockedTicker := <-mockedtimex.NewTickerCalls

			<-mockedTicker.Mock.StoppedChan()

			assert.Panics(t, func() {
				mockedTicker.Mock.Tick(now)
			})
		})
	})
}
