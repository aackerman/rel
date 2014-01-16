package rel

import (
	"bytes"
)

// Used to handle generating Postgres specific sql
type PostgreSQLVisitor struct {
	ToSqlVisitor
}

func (v PostgreSQLVisitor) Accept(visitable Visitable) string {
	return v.Visit(visitable)
}

func (v PostgreSQLVisitor) Visit(visitable Visitable) string {
	switch val := visitable.(type) {
	case *MatchesNode:
		return v.VisitMatchesNode(*val)
	case *DoesNotMatchNode:
		return v.VisitDoesNotMatchNode(*val)
	case *DistinctOnNode:
		return v.VisitDistinctOnNode(*val)
	default:
		return v.ToSqlVisitor.Visit(visitable)
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
