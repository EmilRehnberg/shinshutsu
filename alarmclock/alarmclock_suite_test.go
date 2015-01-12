package alarmclock

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func Test_Alarmclock(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Alarm clock test suite")
}
