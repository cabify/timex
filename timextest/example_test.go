package timextest_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cabify/timex"
	"github.com/cabify/timex/timextest"
)

var now = time.Date(2009, 11, 10, 23, 0, 0, 0, time.UTC)

func ExampleTestImplementation_SetNow() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		mockedtimex.SetNow(time.Unix(1, 0).UTC())
		fmt.Println(timex.Now())
		// Output:
		// 1970-01-01 00:00:01 +0000 UTC
	})
}

func ExampleTestImplementation_Now() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		fmt.Println(timex.Now())
		// Output:
		// 2009-11-10 23:00:00 +0000 UTC
	})
}

func ExampleTestImplementation_Since() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		fmt.Println(timex.Since(now.Add(-time.Hour)))
		// Output:
		// 1h0m0s
	})
}

func ExampleTestImplementation_Until() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		fmt.Println(timex.Until(now.Add(time.Minute)))
		// Output:
		// 1m0s
	})
}

// ExampleTestImplementation_Sleep observes the execution of a tempSwitch that uses timex.Sleep
// To change the state of the program just temporarily
// This way the assertions are deterministic and we don't have to wait the real amount of time
func ExampleTestImplementation_Sleep() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		sw := tempSwitch{new(int64)}
		done := make(chan struct{})

		// Check it's turned off
		fmt.Printf("First state of switch is %t\n", sw.IsTurnedOn())

		go func() {
			sw.TurnOn()
			close(done)
		}()

		// Wait until the program is sleeping
		sleepCall := <-mockedtimex.SleepCalls

		// Check it's turned on
		fmt.Printf("Then it's temporarily %t\n", sw.IsTurnedOn())

		// Let it wake up and wait until TurnOn call finishes execution
		sleepCall.WakeUp()
		<-done

		// Check it's turned off again
		fmt.Printf("And then it's %t again\n", sw.IsTurnedOn())

		// Output:
		// First state of switch is false
		// Then it's temporarily true
		// And then it's false again
	})
}

func ExampleTestImplementation_After() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		var value bool

		wg := sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-timex.After(time.Minute)
			value = false
		}()

		// Note that the race detector would usually complain about this read because it would be
		// racy with the test
		value = true

		(<-mockedtimex.AfterCalls).Mock.Trigger(time.Time{})
		wg.Wait()

		fmt.Println(value)

		// Output:
		// false
	})
}

func ExampleTestImplementation_AfterFunc() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		go func() {
			timex.AfterFunc(time.Second, func() { fmt.Println("This happens a second later") })
			timex.AfterFunc(time.Hour, func() { fmt.Println("This happens an hour later") })
		}()

		firstCall := <-mockedtimex.AfterFuncCalls
		fmt.Printf("First function is scheduled for %s\n", firstCall.Duration)

		secondCall := <-mockedtimex.AfterFuncCalls
		fmt.Printf("Second function is scheduled for %s\n", secondCall.Duration)

		firstCall.Mock.Trigger(now)
		secondCall.Mock.Trigger(now)

		// Output:
		// First function is scheduled for 1s
		// Second function is scheduled for 1h0m0s
		// This happens a second later
		// This happens an hour later
	})
}

func ExampleTestImplementation_NewTicker() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		go func() {
			ticker := timex.NewTicker(time.Hour)
			for t := range ticker.C() {
				fmt.Printf("%s\n", t)
			}
		}()

		tickerCall := <-mockedtimex.NewTickerCalls
		tickerCall.Mock.Tick(now.Add(time.Second))
		tickerCall.Mock.Tick(now.Add(2 * time.Second))

		// Output:
		// 2009-11-10 23:00:01 +0000 UTC
		// 2009-11-10 23:00:02 +0000 UTC
	})
}

func ExampleTestImplementation_NewTimer() {
	timextest.Mocked(now, func(mockedtimex *timextest.TestImplementation) {
		wg := sync.WaitGroup{}
		wg.Add(1)
		go func(d time.Duration) {
			defer wg.Done()
			ticker := timex.NewTimer(d)
			t := <-ticker.C()
			fmt.Println(t)
		}(time.Minute)

		timer := <-mockedtimex.NewTimerCalls
		timer.Mock.Trigger(now.Add(timer.Duration))
		wg.Wait()

		// Output:
		// 2009-11-10 23:01:00 +0000 UTC
	})
}

type tempSwitch struct{ val *int64 }

func (ts tempSwitch) TurnOn() {
	atomic.AddInt64(ts.val, 1)
	timex.Sleep(time.Hour)
	atomic.AddInt64(ts.val, -1)
}

func (ts tempSwitch) IsTurnedOn() bool {
	return atomic.LoadInt64(ts.val) == 1
}
