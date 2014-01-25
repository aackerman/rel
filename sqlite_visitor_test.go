package rel_test

// import (
// 	. "."
// 	. "github.com/onsi/ginkgo"
// 	. "github.com/onsi/gomega"
// )

// var _ = Describe("SQLiteVisitor", func() {
// 	var visitor Visitor

// 	BeforeEach(func() {
// 		Register(NewTestEngine())
// 		visitor = SQLiteVisitor{ToSqlVisitor{Conn: BaseTestConnector{}}}
// 	})

// 	It("defaults limit to -1", func() {
// 		stmt := NewSelectStatementNode()
// 		stmt.Offset = NewOffsetNode(Sql(1))
// 		sql := visitor.Accept(stmt)
// 		Expect(sql).To(Equal("SELECT LIMIT -1 OFFSET 1"))
// 	})

// 	It("does not support locking", func() {
// 		node := &LockNode{Expr: Sql("FOR UPDATE")}
// 		sql := visitor.Accept(node)
// 		Expect(sql).To(Equal(""))
// 	})

// })
