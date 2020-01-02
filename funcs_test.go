// +build !timex_disable

package timex_test

import (
	"testing"
	"time"

	"github.com/cabify/timex"
	"github.com/cabify/timex/timexmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	someDate     = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	someDuration = 288 * time.Second
)

func TestNow(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		mocked.On("Now").Once().Return(someDate)
		defer mocked.AssertExpectations(t)

		now := timex.Now()
		assert.Equal(t, someDate, now)
	})
}

func TestSince(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		mocked.On("Since", someDate).Once().Return(someDuration)
		defer mocked.AssertExpectations(t)

		since := timex.Since(someDate)
		assert.Equal(t, someDuration, since)
	})
}

func TestUntil(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		mocked.On("Until", someDate).Once().Return(someDuration)
		defer mocked.AssertExpectations(t)

		until := timex.Until(someDate)
		assert.Equal(t, someDuration, until)
	})
}

func TestAfterFunc(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		providedCorrectFunction := false
		expectedTimer := &timexmock.Timer{}
		mocked.On("AfterFunc", someDuration, mock.Anything).Once().Return(func(_ time.Duration, f func()) timex.Timer {
			f()
			return expectedTimer
		})
		defer mocked.AssertExpectations(t)

		timer := timex.AfterFunc(someDuration, func() { providedCorrectFunction = true })
		assert.Equal(t, expectedTimer, timer)
		assert.True(t, providedCorrectFunction)
	})
}

func TestSleep(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		mocked.On("Sleep", someDuration).Once()
		defer mocked.AssertExpectations(t)

		timex.Sleep(someDuration)
	})
}

func TestAfter(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		expectedChan := make(<-chan time.Time)
		mocked.On("After", someDuration).Once().Return(expectedChan)
		defer mocked.AssertExpectations(t)

		ch := timex.After(someDuration)
		assert.Equal(t, expectedChan, ch)
	})
}

func TestNewTicker(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		expectedTicker := &timexmock.Ticker{}
		mocked.On("NewTicker", someDuration).Once().Return(expectedTicker)
		defer mocked.AssertExpectations(t)

		ticker := timex.NewTicker(someDuration)
		assert.Equal(t, expectedTicker, ticker)
	})
}

func TestNewTimer(t *testing.T) {
	timexmock.Mocked(func(mocked *timexmock.Implementation) {
		expectedTimer := &timexmock.Timer{}
		mocked.On("NewTimer", someDuration).Once().Return(expectedTimer)
		defer mocked.AssertExpectations(t)

		timer := timex.NewTimer(someDuration)
		assert.Equal(t, expectedTimer, timer)
	})
}
