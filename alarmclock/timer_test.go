package alarmclock

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe(".minutesRemaining", func() {
	Context("given 1 sec", func() {
		It("returns 0", func() {
			duration := time.Second * time.Duration(1)
			Expect(minutesRemaining(duration)).To(Equal(0))
		})
	})

	Context("given 121 sec", func() {
		It("returns 2", func() {
			duration := time.Second * time.Duration(121)
			Expect(minutesRemaining(duration)).To(Equal(2))
		})
	})
})

var _ = Describe(".secondsRemaining", func() {
	Context("given 1 sec", func() {
		It("returns 1", func() {
			duration := time.Second * time.Duration(1)
			Expect(secondsRemaining(duration)).To(Equal(1))
		})
	})

	Context("given 156 sec", func() {
		It("returns 36", func() {
			duration := time.Second * time.Duration(156)
			Expect(secondsRemaining(duration)).To(Equal(36))
		})
	})
})

var _ = Describe(".second", func() {
	It("returns a time.Second", func() {
		Expect(second()).To(Equal(time.Second))
	})
})

var _ = Describe("Timer#duration", func() {
	It("returns the time in seconds that the Timer was initialized with", func() {
		clock := Timer{71}
		Expect(clock.duration()).To(Equal(time.Duration(71) * second()))
	})
})

var _ = Describe(".formattedDuration", func() {
	It("returns seconds in MM:SS format", func() {
		duration := time.Duration(156) * second()
		expectedFormat := "02:36               \r"
		Expect(formattedDuration(duration)).To(Equal(expectedFormat))
	})
})

var _ = Describe("Timer#Countdown", func() {
	var nilFunc func() = func() {}
	Context("on a 71 second timer", func() {
		It("calls ticker#Stop once", func() {
			var stopCalls int
			stopFunc := func() {
				stopCalls += 1
			}
			newTicker = func() ClockTicker {
				return &testTicker{nilFunc, stopFunc}
			}
			Timer{71}.Countdown()
			Expect(stopCalls).To(Equal(1))
		})

		It("calls ticker#Pop 72 times", func() {
			var popCalls int
			popFunc := func() {
				popCalls += 1
			}
			newTicker = func() ClockTicker {
				return &testTicker{popFunc, nilFunc}
			}

			Timer{71}.Countdown()
			Expect(popCalls).To(Equal(72))
		})
	})
})

type testTicker struct {
	popFunc  func()
	stopFunc func()
}

func (t *testTicker) Pop() {
	t.popFunc()
}

func (t *testTicker) Stop() {
	t.stopFunc()
}
