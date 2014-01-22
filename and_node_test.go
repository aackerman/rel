package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AndNode", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("can be equal to other AndNode's", func() {
		and1 := AndNode{Children: &[]Visitable{Sql("foo"), Sql("bar")}}
		and2 := AndNode{Children: &[]Visitable{Sql("foo"), Sql("bar")}}
		Expect(and1.Eq(and2)).To(BeTrue())
	})
})
