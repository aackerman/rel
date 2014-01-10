package rel

import (
	"bytes"
)

type MysqlVisitor struct {
	ToSqlVisitor
}

func (v MysqlVisitor) Accept(a Visitable) string {
	return v.Visit(a)
}

func (v MysqlVisitor) Visit(a Visitable) string {
	switch val := a.(type) {
	case *BinNode:
		return v.VisitBinNode(*val)
	default:
		return v.ToSqlVisitor.Visit(a)
	}
}

func (v MysqlVisitor) VisitBinNode(node BinNode) string {
	var buf bytes.Buffer
	buf.WriteString("BINARY ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}
