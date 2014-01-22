package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InsertStatementNode", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("is equal to other insert statements", func() {
		users1 := NewTable("users")
		users2 := NewTable("users")
		stmt1 := InsertStatementNode{Relation: users1}
		stmt2 := &InsertStatementNode{Relation: users2}
		Expect(stmt1.Eq(stmt2)).To(BeTrue())
	})
})
