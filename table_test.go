package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Table", func() {
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
		Expect(sql).To(Equal(`SELECT FROM "users" "foo" OFFSET 10`))
	})

	It("has an order method", func() {
		table := NewTable("users")
		sm := table.Order(Sql("foo"))
		sql := sm.ToSql()
		Expect(sql).To(Equal(`SELECT FROM "users" ORDER BY foo`))
	})

	It("has a take method", func() {
		table := NewTable("users")
		sm := table.Take(1)
		sm.Project(Sql("*"))
		sql := sm.ToSql()
		Expect(sql).To(Equal(`SELECT * FROM "users" LIMIT 1`))
	})

	It("has a where method", func() {
		table := NewTable("users")
		sm := table.Where(table.Attr("id").Eq(Sql(1)))
		sm.Project(table.Attr("id"))
		sql := sm.ToSql()
		Expect(sql).To(Equal(`SELECT "users"."id" FROM "users" WHERE "users"."id" = 1`))
	})

	It("has a project method", func() {
		table := NewTable("users")
		query := table.Project(Sql("*"))
		sql := query.ToSql()
		Expect(sql).To(Equal(`SELECT * FROM "users"`))
	})

	It("has a skip method", func() {
		table := NewTable("users")
		query := table.Skip(2)
		sql := query.ToSql()
		Expect(sql).To(Equal(`SELECT FROM "users" OFFSET 2`))
	})

	It("has an offset method", func() {
		table := NewTable("users")
		query := table.Offset(2)
		sql := query.ToSql()
		Expect(sql).To(Equal(`SELECT FROM "users" OFFSET 2`))
	})

	It("has a having method", func() {
		table := NewTable("users")
		query := table.Having(table.Attr("id").Eq(Sql(10)))
		sql := query.ToSql()
		Expect(sql).To(Equal(`SELECT FROM "users" HAVING "users"."id" = 10`))
	})

	It("has a group method", func() {
		table := NewTable("users")
		query := table.Group(table.Attr("id"))
		sql := query.ToSql()
		Expect(sql).To(Equal(`SELECT FROM "users" GROUP BY "users"."id"`))
	})

	It("Table#Project accepts multiple arguments", func() {
		table := NewTable("users")
		query := table.Project(Sql("*"), Sql("*"))
		sql := query.ToSql()
		Expect(sql).To(Equal(`SELECT *, * FROM "users"`))
	})

	It("can return a selectmanager", func() {
		table := NewTable("")
		sm := table.From(table)
		sql := sm.ToSql()
		Expect(sql).To(Equal(`SELECT`))
	})
})
