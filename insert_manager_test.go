package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("InsertManager", func() {
	BeforeEach(func() {
		Register("postgresql", NewEngine())
	})

	It("can create values", func() {
		mgr := NewInsertManager(DefaultEngine)
		values := make([]interface{}, 0)
		values = append(values, Star())
		columns := []AttributeNode{AttributeNode{}}
		mgr.SetValues(mgr.CreateValues(values, columns))
		sql := mgr.ToSql()
		expected := "INSERT INTO NULL VALUES (*)"
		Expect(sql).To(Equal(expected))
	})

	It("can insert values", func() {
		users := NewTable("users")
		mgr := NewInsertManager(DefaultEngine)
		mgr.Ast.Relation = users
		mgr.Insert(users.Attr("bool"), false)
		sql := mgr.ToSql()
		expected := "INSERT INTO \"users\" (\"bool\") VALUES ('f')"
		Expect(sql).To(Equal(expected))
	})
})
