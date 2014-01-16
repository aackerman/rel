package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AscendingNode", func() {
	BeforeEach(func() {
		Register("postgresql", NewEngine())
	})

	It("can be equal to another AscendingNode", func() {
		asc1 := AscendingNode{Expr: Sql("zomg")}
		asc2 := AscendingNode{Expr: Sql("zomg")}
		Expect(asc1.Eq(asc2)).To(BeTrue())
	})

	It("returns it's direction", func() {
		asc := AscendingNode{Expr: Sql("zomg")}
		Expect(asc.Direction()).To(Equal("ASC"))
	})

	It("can return a reversed node", func() {
		asc := AscendingNode{Expr: Sql("zomg")}
		desc := asc.Reverse()
		Expect(desc.Expr).To(Equal(Sql("zomg")))
		Expect(desc.Direction()).To(Equal("DESC"))
	})
})
