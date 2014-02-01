package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateManager", func() {
	It("handles limit properly", func() {
		table := NewTable("users")
		mgr := NewUpdateManager(RelEngine)
		mgr.SetTable(table)
		mgr.Set(table.Attr("name"), NewBindParamNode("?"))
		Expect(mgr.ToSql()).To(Equal("UPDATE \"users\" SET \"name\" = ?"))
	})
})
