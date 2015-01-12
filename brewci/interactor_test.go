package brewci

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	teas map[string][]int = map[string][]int{
		"nisemono-cha": []int{2, 3},
	}
)

var _ = Describe("brewci package", func() {
	Describe("#parseInt", func() {
		It("returns a 3 when given '3'", func() {
			Expect(parseInt('3')).To(Equal(3))
		})
	})

	Describe("#Execute", func() {
		BeforeEach(func() {
			teaNames = []string{}
			p = fmt.Scan
			puts = p
			printf = fmt.Scanf
		})

		Context("given the first steep time (out of multiple), and then stopping", func() {
			BeforeEach(func() {
				teaQueryResponse = func() rune { return '0' }
				steepNumberQueryResponse = func() rune { return '1' }

				runTimer = func(seconds int) {}
				continueResponse = func() rune { return 's' }
			})

			It("queries for which tea to pick once", func() {
				var teaQueryNumber int
				teaQueryResponse = func() rune {
					teaQueryNumber += 1
					return '0'
				}

				Execute(teas)
				Expect(teaQueryNumber).To(Equal(1))
			})

			It("queries for steep number once", func() {
				var steepNumberQueryNumber int
				steepNumberQueryResponse = func() rune {
					steepNumberQueryNumber += 1
					return '1'
				}

				Execute(teas)
				Expect(steepNumberQueryNumber).To(Equal(1))
			})

			It("queries for continue once", func() {
				var continutationQueryNumber int
				continueResponse = func() rune {
					continutationQueryNumber += 1
					return 's'
				}

				Execute(teas)
				Expect(continutationQueryNumber).To(Equal(1))
			})

			It("calls runTimer once", func() {
				var runTimerCallNumber int
				runTimer = func(seconds int) {
					runTimerCallNumber += 1
				}

				Execute(teas)
				Expect(runTimerCallNumber).To(Equal(1))
			})

			It("calls runTimer with 2", func() {
				var calledRunTimerWith int
				runTimer = func(seconds int) {
					calledRunTimerWith = seconds
				}

				Execute(teas)
				Expect(calledRunTimerWith).To(Equal(2))
			})
		})

		Context("given the first steep time (out of two), and then continuing", func() {
			BeforeEach(func() {
				teaQueryResponse = func() rune { return '0' }
				steepNumberQueryResponse = func() rune { return '1' }

				runTimer = func(seconds int) {}
				continueResponse = func() rune { return 'c' }

				teas = map[string][]int{
					"nisemono-cha": []int{2, 3},
				}
			})

			It("runs the timer twice", func() {
				var runTimerCallNumber int
				runTimer = func(seconds int) {
					runTimerCallNumber += 1
				}

				Execute(teas)
				Expect(runTimerCallNumber).To(Equal(2))
			})

			It("runs the timer with the brewing times in order", func() {
				var runTimerSecCalls []int
				runTimer = func(seconds int) {
					runTimerSecCalls = append(runTimerSecCalls, seconds)
				}

				Execute(teas)
				Expect(runTimerSecCalls).To(Equal([]int{2, 3}))
			})

			It("queries for which tea to pick once", func() {
				var teaQueryNumber int
				teaQueryResponse = func() rune {
					teaQueryNumber += 1
					return '0'
				}

				Execute(teas)
				Expect(teaQueryNumber).To(Equal(1))
			})

			It("queries for steep number once", func() {
				var steepNumberQueryNumber int
				steepNumberQueryResponse = func() rune {
					steepNumberQueryNumber += 1
					return '1'
				}

				Execute(teas)
				Expect(steepNumberQueryNumber).To(Equal(1))
			})

			It("queries for continue once", func() {
				var continutationQueryNumber int
				continueResponse = func() rune {
					continutationQueryNumber += 1
					return 'c'
				}

				Execute(teas)
				Expect(continutationQueryNumber).To(Equal(1))
			})
		})

		Context("given the second steep time (out of three), and then continuing", func() {
			BeforeEach(func() {
				teas = map[string][]int{
					"nisemono-cha": []int{2, 3, 4},
				}
				teaQueryResponse = func() rune { return '0' }
				steepNumberQueryResponse = func() rune { return '2' }

				runTimer = func(seconds int) {}
				continueResponse = func() rune { return 'c' }
			})

			It("queries for steep number once", func() {
				var steepNumberQueryNumber int
				steepNumberQueryResponse = func() rune {
					steepNumberQueryNumber += 1
					return '1'
				}

				Execute(teas)
				Expect(steepNumberQueryNumber).To(Equal(1))
			})

			It("queries for continue once", func() {
				var continutationQueryNumber int
				continueResponse = func() rune {
					continutationQueryNumber += 1
					return 'c'
				}

				Execute(teas)
				Expect(continutationQueryNumber).To(Equal(1))
			})

			It("runs the timer with the brewing times in order", func() {
				var runTimerSecCalls []int
				runTimer = func(seconds int) {
					runTimerSecCalls = append(runTimerSecCalls, seconds)
				}

				Execute(teas)
				Expect(runTimerSecCalls).To(Equal([]int{3, 4}))
			})
		})

		Context("given the NO steep time, and then stopping", func() {
			BeforeEach(func() {
				teaQueryResponse = func() rune { return '0' }
				steepNumberQueryResponse = func() rune { return '\n' }

				runTimer = func(seconds int) {}
				continueResponse = func() rune { return 's' }
			})

			It("queries for steep number once", func() {
				var steepNumberQueryNumber int
				steepNumberQueryResponse = func() rune {
					steepNumberQueryNumber += 1
					return '\n'
				}

				Execute(teas)
				Expect(steepNumberQueryNumber).To(Equal(1))
			})

			It("calls runTimer once", func() {
				var runTimerCallNumber int
				runTimer = func(seconds int) {
					runTimerCallNumber += 1
				}

				Execute(teas)
				Expect(runTimerCallNumber).To(Equal(1))
			})

			It("calls runTimer with 2", func() {
				var calledRunTimerWith int
				runTimer = func(seconds int) {
					calledRunTimerWith = seconds
				}

				Execute(teas)
				Expect(calledRunTimerWith).To(Equal(2))
			})
		})
	})

	Describe("#userQuits", func() {
		It("calls #continueResponse", func() {
			var continueResponseWasCalled bool
			continueResponse = func() rune {
				continueResponseWasCalled = true
				return 'c'
			}
			userQuits()

			Expect(continueResponseWasCalled).To(Equal(true))
		})

		Context("when the user returns a 'c'", func() {
			It("returns false", func() {
				continueResponse = func() rune {
					return 'c'
				}
				toQuit := userQuits()

				Expect(toQuit).To(Equal(false))
			})
		})

		Context("when the user returns nothing", func() {
			It("returns false", func() {
				continueResponse = func() rune {
					return '\n'
				}
				toQuit := userQuits()

				Expect(toQuit).To(Equal(false))
			})
		})

		Context("when the user returns a 's'", func() {
			It("returns true", func() {
				continueResponse = func() rune {
					return 's'
				}
				toQuit := userQuits()
				Expect(toQuit).To(Equal(true))
			})
		})

		Context("when the user returns a 'a' then a 's'", func() {
			var (
				continueResponseCallTimes int
				toQuit                    bool
			)

			BeforeEach(func() {
				continueResponseCallTimes = 0
				continueResponse = func() (reply rune) {
					if continueResponseCallTimes > 0 {
						reply = 's'
					} else {
						reply = 'a'
					}
					continueResponseCallTimes += 1
					return
				}
				toQuit = userQuits()
			})

			It("returns true", func() {
				Expect(toQuit).To(Equal(true))
			})

			It("calls #continueResponse twice", func() {
				Expect(continueResponseCallTimes).To(Equal(2))
			})
		})
	}) // ends "#userQuits" description
})
