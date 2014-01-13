package rel

import (
	"bytes"
)

type PostgreSQLVisitor struct {
	ToSqlVisitor
}

func (v PostgreSQLVisitor) Accept(a Visitable) string {
	return v.Visit(a)
}

func (v PostgreSQLVisitor) Visit(a Visitable) string {
	switch val := a.(type) {
	case *MatchesNode:
		return v.VisitMatchesNode(*val)
	case *DoesNotMatchNode:
		return v.VisitDoesNotMatchNode(*val)
	case *DistinctOnNode:
		return v.VisitDistinctOnNode(*val)
	default:
		return v.ToSqlVisitor.Visit(a)
	}
}

func (v PostgreSQLVisitor) VisitMatchesNode(node MatchesNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" ILIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v PostgreSQLVisitor) VisitDoesNotMatchNode(node DoesNotMatchNode) string {
	var buf bytes.Buffer
	buf.WriteString(v.Visit(node.Left))
	buf.WriteString(" NOT ILIKE ")
	buf.WriteString(v.Visit(node.Right))
	return buf.String()
}

func (v PostgreSQLVisitor) VisitDistinctOnNode(node DistinctOnNode) string {
	var buf bytes.Buffer
	buf.WriteString("DISTINCT ON ( ")
	buf.WriteString(v.Visit(node.Expr))
	buf.WriteString(" )")
	return buf.String()
}
