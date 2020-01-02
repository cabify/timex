package timextest_test

import (
	"testing"
	"time"

	"github.com/cabify/timex"
	"github.com/cabify/timex/timextest"
	"github.com/stretchr/testify/assert"
)

func TestMockedTimer_StoppedChan(t *testing.T) {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		stopValue := make(chan bool)
		go func() {
			ticker := timex.NewTimer(time.Second)
			stopValue <- ticker.Stop()
		}()
		mockedTicker := <-mockedtimex.NewTimerCalls
		mockedTicker.Mock.StopValue(true)

		select {
		case <-mockedTicker.Mock.StoppedChan():
		case <-time.After(time.Second):
			t.Errorf("Stop should have been called")
		}

		assert.True(t, <-stopValue)
	})
}

func TestMockedTimer_Trigger(t *testing.T) {
	t.Run("panics on stopped timer", func(t *testing.T) {
		timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
			go func() {
				ticker := timex.NewTimer(time.Second)
				ticker.Stop()
			}()
			mockedTimer := <-mockedtimex.NewTimerCalls
			mockedTimer.Mock.StopValue(true)

			<-mockedTimer.Mock.StoppedChan()

			assert.Panics(t, func() {
				mockedTimer.Mock.Trigger(now)
			})
		})
	})
}
