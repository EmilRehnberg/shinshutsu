package alarmclock

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"strings"
)

var _ = Describe(".Toll", func() {
	Context("given a linux system", func() {
		BeforeEach(func() {
			systemOs = func() string { return "linux" }
		})

		It("tolls a bell once", func() {
			var runCalls int

			cmd := testCmd{
				runFunc: func() {
					runCalls += 1
				},
			}

			buildCommand = func(exe, path string) CmdRunner {
				return &cmd
			}

			Toll()
			Expect(runCalls).To(Equal(1))
		})

		It("tolls the bell with paplay", func() {
			var bellExecutable string

			cmd := testCmd{runFunc: func() {}}

			buildCommand = func(exe, path string) CmdRunner {
				bellExecutable = exe
				return &cmd
			}

			Toll()
			Expect(bellExecutable).To(Equal("paplay"))
		})

		It("tolls the bell with sound file path", func() {
			var bellSoundPath string

			cmd := testCmd{runFunc: func() {}}

			buildCommand = func(exe, path string) CmdRunner {
				bellSoundPath = path
				return &cmd
			}

			Toll()
			Expect(bellSoundPath).To(Equal("/usr/share/sounds/freedesktop/stereo/alarm-clock-elapsed.oga"))
		})
	})

	Context("given a osx system", func() {
		BeforeEach(func() {
			systemOs = func() string { return "darwin" }
		})

		It("tolls a bell once", func() {
			var runCalls int

			cmd := testCmd{
				runFunc: func() {
					runCalls += 1
				},
			}

			buildCommand = func(exe, path string) CmdRunner {
				return &cmd
			}

			Toll()
			Expect(runCalls).To(Equal(1))
		})

		It("tolls the bell with afplay", func() {
			var bellExecutable string

			cmd := testCmd{runFunc: func() {}}

			buildCommand = func(exe, path string) CmdRunner {
				bellExecutable = exe
				return &cmd
			}

			Toll()
			Expect(bellExecutable).To(Equal("afplay"))
		})

		It("tolls the bell with the repo media file", func() {
			var bellSoundPath string

			cmd := testCmd{runFunc: func() {}}

			buildCommand = func(exe, path string) CmdRunner {
				bellSoundPath = path
				return &cmd
			}

			Toll()
			Expect(strings.HasSuffix(bellSoundPath, "github.com/emilrehnberg/shinshutsu/media/alarm-clock.aiff")).To(Equal(true))
		})
	})

	Context("given a unknown system (e.g. Windows)", func() {
		BeforeEach(func() {
			systemOs = func() string { return "windows" }
			printf = fmt.Scanf
			dirtyExit = func() {}
		})

		It("does not toll a bell", func() {
			var runCalls int

			cmd := testCmd{
				runFunc: func() {
					runCalls += 1
				},
			}

			buildCommand = func(exe, path string) CmdRunner {
				return &cmd
			}

			Toll()
			Expect(runCalls).To(Equal(0))
		})

		It("exists dirty", func() {
			var dirtyExitCalls int
			dirtyExit = func() {
				dirtyExitCalls += 1
			}
			Toll()
			Expect(dirtyExitCalls).To(Equal(1))
		})
	})
})

type testCmd struct {
	runFunc func()
}

func (t *testCmd) Run() error {
	t.runFunc()
	return nil
}
