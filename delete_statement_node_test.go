package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeleteStatementNode", func() {
	BeforeEach(func() {
		Register("postgresql", NewEngine())
	})

	It("can be equal to other DeleteStatementNode's", func() {
		ds1 := DeleteStatementNode{
			Wheres: &[]Visitable{Sql("a"), Sql("b"), Sql("c")},
		}

		ds2 := DeleteStatementNode{
			Wheres: &[]Visitable{Sql("a"), Sql("b"), Sql("c")},
		}
		Expect(ds1.Eq(ds2)).To(BeTrue())
	})
})
