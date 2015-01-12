package brewci

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestBrewci(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Brewci Suite")
}
