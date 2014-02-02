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
		mgr.SetKey(Sql("id"))
		mgr.Take(10)
		mgr.Set(table.Attr("name"), nil)
		Expect(mgr.ToSql()).To(Equal("UPDATE \"users\" SET \"name\" = ?"))
	})
})
