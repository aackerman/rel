package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateManager", func() {
	It("should not quote bind params", func() {
		table := NewTable("users")
		mgr := NewUpdateManager(RelEngine)
		mgr.SetTable(table)
		mgr.Set(table.Attr("name"), NewBindParamNode("?"))
		Expect(mgr.ToSql()).To(Equal("UPDATE \"users\" SET \"name\" = ?"))
	})

	It("should handle limit properly", func() {
		table := NewTable("users")
		mgr := NewUpdateManager(RelEngine)
		mgr.SetTable(table)
		mgr.SetKey(table.Attr("id"))
		mgr.Take(10)
		mgr.Set(table.Attr("name"), nil)
		Expect(mgr.ToSql()).To(Equal("UPDATE \"users\" SET \"name\" = NULL WHERE \"users\".\"id\" IN (SELECT \"users\".\"id\" FROM \"users\" LIMIT 10)"))
	})

	It("updates with null", func() {
		table := NewTable("users")
		mgr := NewUpdateManager(RelEngine)
		mgr.SetTable(table)
		mgr.Set(table.Attr("name"), nil)
		Expect(mgr.ToSql()).To(Equal("UPDATE \"users\" SET \"name\" = NULL"))
	})

	Describe("SetTable", func() {
		It("sets the relation", func() {
			mgr := NewUpdateManager(RelEngine)
			mgr.SetTable(NewTable("users"))
			Expect(mgr.ToSql()).To(Equal("UPDATE \"users\""))
		})
	})

	Describe("Where", func() {
		It("generates a where clause", func() {
			table := NewTable("users")
			mgr := NewUpdateManager(RelEngine)
			mgr.SetTable(table)
			mgr.Where(table.Attr("id").Eq(Sql(1)))
			Expect(mgr.ToSql()).To(Equal("UPDATE \"users\" WHERE \"users\".\"id\" = 1"))
		})
	})
})
