package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Table", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("has a name", func() {
		table := NewTable("users")
		Expect(table.Name).To(Equal("users"))
	})

	It("has a table alias", func() {
		table := NewTable("users")
		alias := table.Alias()
		Expect(alias.Name.Raw).To(Equal("users_2"))
	})

	It("can set the table alias", func() {
		table := NewTable("users")
		table.SetTableAlias("foo")
		manager := table.From(table)
		manager.Skip(10)
		sql := manager.ToSql()
		expected := "SELECT FROM \"users\" \"foo\" OFFSET 10"
		Expect(sql).To(Equal(expected))
	})

	It("has an order method", func() {
		table := NewTable("users")
		sm := table.Order(Sql("foo"))
		sql := sm.ToSql()
		expected := "SELECT FROM \"users\" ORDER BY foo"
		Expect(sql).To(Equal(expected))
	})

	It("has a take method", func() {
		table := NewTable("users")
		sm := table.Take(1)
		sm.Project(Sql("*"))
		sql := sm.ToSql()
		expected := "SELECT * FROM \"users\" LIMIT 1"
		Expect(sql).To(Equal(expected))
	})

	It("has a where method", func() {
		table := NewTable("users")
		sm := table.Where(table.Attr("id").Eq(Sql(1)))
		sm.Project(table.Attr("id"))
		sql := sm.ToSql()
		expected := "SELECT \"users\".\"id\" FROM \"users\" WHERE \"users\".\"id\" = 1"
		Expect(sql).To(Equal(expected))
	})

	It("has a project method", func() {
		table := NewTable("users")
		query := table.Project(Sql("*"))
		sql := query.ToSql()
		expected := "SELECT * FROM \"users\""
		Expect(sql).To(Equal(expected))
	})

	It("has a skip method", func() {
		table := NewTable("users")
		query := table.Skip(2)
		sql := query.ToSql()
		expected := "SELECT FROM \"users\" OFFSET 2"
		Expect(sql).To(Equal(expected))
	})

	It("has an offset method", func() {
		table := NewTable("users")
		query := table.Offset(2)
		sql := query.ToSql()
		expected := "SELECT FROM \"users\" OFFSET 2"
		Expect(sql).To(Equal(expected))
	})

	It("has a having method", func() {
		table := NewTable("users")
		query := table.Having(table.Attr("id").Eq(Sql(10)))
		sql := query.ToSql()
		expected := "SELECT FROM \"users\" HAVING \"users\".\"id\" = 10"
		Expect(sql).To(Equal(expected))
	})

	It("has a group method", func() {
		table := NewTable("users")
		query := table.Group(table.Attr("id"))
		sql := query.ToSql()
		expected := "SELECT FROM \"users\" GROUP BY \"users\".\"id\""
		Expect(sql).To(Equal(expected))
	})

	It("Table#Project accepts multiple arguments", func() {
		table := NewTable("users")
		query := table.Project(Sql("*"), Sql("*"))
		sql := query.ToSql()
		expected := "SELECT *, * FROM \"users\""
		Expect(sql).To(Equal(expected))
	})

	It("can return a selectmanager", func() {
		table := NewTable("")
		sm := table.From(table)
		sql := sm.ToSql()
		expected := "SELECT"
		Expect(sql).To(Equal(expected))
	})
})
