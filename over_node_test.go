package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("AscendingNode", func() {
	BeforeEach(func() {
		Register("postgresql", NewEngine())
	})

	It("should alias the expression", func() {
		users := NewTable("users")
		mgr := users.From(users)
		mgr.Project(users.Attr("id").Count().Over(nil).As(Sql("foo")))
		sql := mgr.ToSql()
		expected := "SELECT COUNT(\"users\".\"id\") OVER () AS foo FROM \"users\""
		Expect(sql).To(Equal(expected))
	})
})
