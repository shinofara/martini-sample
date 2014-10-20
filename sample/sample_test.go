package sample_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sample", func() {
	It("ginkgoを使ったサンプルテスト", func() {
		a := "test"
	  b	 := "test"

		//Option+z -> Ω
		Ω(a).Should(Equal(b))
		Expect(a).Should(Equal(b))
	})

})
