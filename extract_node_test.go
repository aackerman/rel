package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ExtractNode", func() {
	It("should extract field", func() {
		table := NewTable("users")
		mgr := table.From(table)
		mgr.Project(table.Attr("timestamp").Extract(Sql("date")))
		Expect(mgr.ToSql()).To(Equal("SELECT EXTRACT(DATE FROM \"users\".\"timestamp\") FROM \"users\""))
	})

	Describe("As", func() {
		It("should alias the Extract", func() {
			table := NewTable("users")
			mgr := table.From(table)
			mgr.Project(table.Attr("timestamp").Extract(Sql("date")).As(Sql("foo")))
			Expect(mgr.ToSql()).To(Equal("SELECT EXTRACT(DATE FROM \"users\".\"timestamp\") AS foo FROM \"users\""))
		})
	})
})
