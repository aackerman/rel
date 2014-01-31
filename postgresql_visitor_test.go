package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PostgreSQLVisitor", func() {
	var visitor Visitor

	BeforeEach(func() {
		visitor = &PostgreSQLVisitor{Conn: DefaultConnector{}}
	})

	It("should support distinct on", func() {
		core := NewSelectCoreNode()
		core.SetQuantifier = NewDistinctOnNode(Sql("aaron"))
		Expect(visitor.Accept(core)).To(Equal("SELECT DISTINCT ON ( aaron )"))
	})

	It("should support distinct", func() {
		core := NewSelectCoreNode()
		core.SetQuantifier = &DistinctNode{}
		Expect(visitor.Accept(core)).To(Equal("SELECT DISTINCT"))
	})

})
