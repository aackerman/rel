package rel

type SQLiteVisitor struct {
	ToSqlVisitor
}

func (v SQLiteVisitor) Accept(visitable Visitable) string {
	return v.Visit(visitable)
}

func (v SQLiteVisitor) Visit(visitable Visitable) string {
	switch val := visitable.(type) {
	case *LockNode:
		return v.VisitLockNode(val)
	case *SelectStatementNode:
		return v.VisitSelectStatementNode(val)
	default:
		return v.ToSqlVisitor.Visit(visitable)
	}
}

// VisitLockNode is overwritten for the SQLiteVisitor
// Locks are not supported in SQLite
func (v SQLiteVisitor) VisitLockNode(node *LockNode) string {
	return ""
}

func (v SQLiteVisitor) VisitSelectStatementNode(node *SelectStatementNode) string {
	if node.Offset != nil && node.Limit == nil {
		node.Limit = &LimitNode{Expr: Sql("-1")}
	}
	return v.ToSqlVisitor.VisitSelectStatementNode(node)
}
