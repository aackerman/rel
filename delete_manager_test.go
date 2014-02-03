package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DeleteManager", func() {
	Describe("From", func() {
		It("sets the relation", func() {
			mgr := NewDeleteManager(RelEngine)
			mgr.From(NewTable("users"))
			Expect(mgr.ToSql()).To(Equal("DELETE FROM \"users\""))
		})
	})

	Describe("Where", func() {
		It("generates a where clause", func() {
			table := NewTable("users")
			mgr := NewDeleteManager(RelEngine)
			mgr.From(table)
			mgr.Where(table.Attr("id").Eq(Sql(1)))
			Expect(mgr.ToSql()).To(Equal("DELETE FROM \"users\" WHERE \"users\".\"id\" = 1"))
		})
	})
})
