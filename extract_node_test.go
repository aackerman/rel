package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	// "testing"
)

var _ = Describe("ExtractNode", func() {
	BeforeEach(func() {
		Register("postgresql", NewEngine())
	})

	It("should extract field", func() {
		table := NewTable("users")
		mgr := table.From(table)
		mgr.Project(table.Attr("timestamp").Extract(Sql("date")))
		Expect(mgr.ToSql()).To(Equal("SELECT EXTRACT(DATE FROM \"users\".\"timestamp\") FROM \"users\""))
	})
})
