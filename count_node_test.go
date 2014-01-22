package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CountNode", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("can be equal to other CountNode's", func() {
		count1 := CountNode{Expressions: []Visitable{Sql("foo")}}
		count2 := CountNode{Expressions: []Visitable{Sql("foo")}}
		Expect(count1.Eq(count2)).To(BeTrue())
	})

	It("can use the As predication to be aliased", func() {
		table := NewTable("users")
		mgr := table.Select(table.Attr("id").Count().As(Sql("foo")))
		sql := mgr.ToSql()
		expected := "SELECT COUNT(\"users\".\"id\") AS foo FROM \"users\""
		Expect(sql).To(Equal(expected))
	})
})
