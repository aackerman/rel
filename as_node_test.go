package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AsNode", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("can be equal to other AsNode's", func() {
		as1 := AsNode{Left: Sql("foo"), Right: Sql("bar")}
		as2 := AsNode{Left: Sql("foo"), Right: Sql("bar")}
		Expect(as1.Eq(as2)).To(BeTrue())
	})
})
