package brewci

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BuildTeasMap", func() {
	It("it builds a tea map", func() {
		getBrewingDataFilePath = func() string { return "./test-data/test_tea_data.yml" }
		expectedTeas := map[string][]int{"nisemono-cha": []int{2, 1, 3}}
		Expect(BuildTeasMap()).To(Equal(expectedTeas))
	})
})
