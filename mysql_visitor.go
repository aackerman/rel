package rel

import (
	"bytes"
)

type MysqlVisitor struct {
	ToSqlVisitor
}

func (v MysqlVisitor) Accept(visitable Visitable) string {
	return v.Visit(visitable)
}

func (v MysqlVisitor) Visit(visitable Visitable) string {
	switch val := visitable.(type) {
	case *BinNode:
		return v.VisitBinNode(*val)
	default:
		return v.ToSqlVisitor.Visit(visitable)
	}
}

func (v MysqlVisitor) VisitBinNode(node BinNode) string {
	var buf bytes.Buffer
	buf.WriteString("BINARY ")
	buf.WriteString(v.Visit(node.Expr))
	return buf.String()
}
