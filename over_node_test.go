package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AscendingNode", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("should alias the expression", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Project(users.Attr("id").Count().Over(nil).As(Sql("foo")))
		sql := mgr.ToSql()
		expected := "SELECT COUNT(\"users\".\"id\") OVER () AS foo FROM \"users\""
		Expect(sql).To(Equal(expected))
	})

	It("should reference the window definition by name", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Project(users.Attr("id").Count().Over(Sql("foo")))
		sql := mgr.ToSql()
		expected := "SELECT COUNT(\"users\".\"id\") OVER foo FROM \"users\""
		Expect(sql).To(Equal(expected))
	})

	It("should use definition in sub-expression", func() {
		users := NewTable("users")
		mgr := users.From(users)
		window := (&WindowNode{}).Order(users.Attr("foo"))
		mgr.Project(users.Attr("id").Count().Over(window))
		sql := mgr.ToSql()
		expected := "SELECT COUNT(\"users\".\"id\") OVER (ORDER BY \"users\".\"foo\") FROM \"users\""
		Expect(sql).To(Equal(expected))
	})
})
