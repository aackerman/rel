package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DescendingNode", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("can be equal to other DescendingNode's", func() {
		desc1 := DescendingNode{Expr: Sql("zomg")}
		desc2 := DescendingNode{Expr: Sql("zomg")}
		Expect(desc1.Eq(desc2)).To(BeTrue())
	})

	It("it has a direction", func() {
		desc := DescendingNode{Expr: Sql("zomg")}
		Expect(desc.Direction()).To(Equal("DESC"))
	})

	It("can reverse direction", func() {
		desc := DescendingNode{Expr: Sql("zomg")}
		asc := desc.Reverse()
		Expect(asc.Direction()).To(Equal("ASC"))
		Expect(asc.Expr).To(Equal(Sql("zomg")))
	})
})
