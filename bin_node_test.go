package rel_test

import (
	. "."
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BinNode", func() {
	BeforeEach(func() {
		Register(NewTestEngine())
	})

	It("can be Equal to another BinNode", func() {
		bin1 := BinNode{Expr: Sql("zomg")}
		bin2 := BinNode{Expr: Sql("zomg")}
		Expect(bin1).To(Equal(bin2))
	})

	It("is visited differently using the MysqlVisitor", func() {
		viz := MysqlVisitor{ToSqlVisitor{Conn: BaseTestConnector{}}}
		bin := &BinNode{Expr: Sql("zomg")}
		sql := viz.Accept(bin)
		Expect(sql).To(Equal("BINARY zomg"))
	})
})
