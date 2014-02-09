package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InsertManager", func() {
	It("can insert values", func() {
		users := NewTable("users")
		mgr := NewInsertManager(RelEngine)
		mgr.Into(users)
		mgr.Insert(users.Attr("email"), Sql("a@b.com"))
		sql := mgr.ToSql()
		expected := `INSERT INTO "users" ("email") VALUES ('a@b.com')`
		Expect(sql).To(Equal(expected))
	})
})
